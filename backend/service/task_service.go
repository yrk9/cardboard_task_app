package service

import (
	"time"
	"database/sql"
	"errors"
	"github.com/yrk9/cardboard_task_app/model"
	"github.com/yrk9/cardboard_task_app/repository"
)

var (
	ErrTaskNotFound = errors.New("タスクが見つかりません")
	ErrTitleRequired   = errors.New("タイトルは必須です")
    ErrInvalidPriority = errors.New("優先度が不正です")
)

type TaskInput struct {
	Title string
	Description string
	Priority string
	DueDate *time.Time
}

type TaskService interface {
	List(userID int64) ([]*model.Task, error)
	Get(userID int64, taskID int64) (*model.Task, error)
	Create(userID int64, input TaskInput) error
	Delete(userID int64, taskID int64) error
	Update(userID int64, taskID int64, input TaskInput) (*model.Task, error)
	ToggleComplete(userID int64, taskID int64) (*model.Task, error)
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (t* taskService) List(userID int64) ([]*model.Task, error) {
	taskList, err := t.taskRepo.FindByUserID(userID)

	if err != nil {
		return nil, err
	}
	return taskList, nil
}

func (t* taskService) Get(userID int64, taskID int64) (*model.Task, error) {
	task, err := t.taskRepo.FindByID(taskID, userID)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrTaskNotFound 
	}

	if err != nil {
		return nil, err
	}

	return task, nil 
}

func (t* taskService) Create(userID int64, input TaskInput) error {
	//バリデーション
	if input.Title == "" {
		return ErrTitleRequired
	}
	if input.Priority != "high" && input.Priority != "mid" && input.Priority != "low" {
		return ErrInvalidPriority
	}
	
	task := &model.Task{
		UserID: userID,
		Title: input.Title,
		Description: input.Description,
		Priority: input.Priority,
		DueDate: input.DueDate,
	}
		
	if err := t.taskRepo.Create(task); err != nil {
		return err
	}

	return nil
}


func (t* taskService) Delete(userID int64, taskID int64) error {
	if err := t.taskRepo.Delete(taskID, userID); err != nil {
		return err
	}

	return nil
}

func (t* taskService) Update(userID int64, taskID int64, input TaskInput) (*model.Task, error) {
	// ① まず既存タスクを取得（自分のタスクか確認も兼ねる）
	task, err := t.taskRepo.FindByID(taskID, userID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrTaskNotFound
	}
	if err != nil {
		return nil, err
	}

	// ② バリデーション
	if input.Title == "" {
		return nil, ErrTitleRequired
	}
	if input.Priority != "high" && input.Priority != "mid" && input.Priority != "low" {
		return nil, ErrInvalidPriority
	}

	// ③ 値を書き換える
	task.Title = input.Title
	task.Description = input.Description
	task.Priority = input.Priority
	task.DueDate = input.DueDate

	// ④ 保存
	if err := t.taskRepo.Update(task); err != nil {
		return nil, err
	}
	return task, nil   // ← すでに手元にある task をそのまま返せる
}

func (t* taskService) ToggleComplete(userID int64, taskID int64) (*model.Task, error) {
	//1件取得
	task, err := t.Get(userID, taskID)

	if err != nil {
		return nil, err
	}

	if task.CompletedAt == nil {
		task = &model.Task{
			CompletedAt: nil,
		}
	} else {
		now := time.Now()
		task = &model.Task{
			CompletedAt: &now,
		}
	}
	return task, nil
}
