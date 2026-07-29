package repository

import (
	"database/sql"
	"time"
	"github.com/yrk9/cardboard_task_app/model"
)

type TaskRepository interface {
	Create(task *model.Task) error
	FindByUserID(userID int64) ([]*model.Task, error)
	FindByID(id int64, userID int64) (*model.Task, error)
	Update(task *model.Task) error
	Delete(id int64, userID int64) error
}

type taskRepository struct {
	db *sql.DB
}
//コンストラクタ
func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

//タスクの作成
func (r taskRepository) Create(task *model.Task) error {
	result, err := r.db.Exec(
		`INSERT INTO tasks (user_id, title, description, priority, due_date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		task.UserID, 
		task.Title, 
		task.Description,
		task.Priority,
		task.DueDate,
		time.Now(), 
		time.Now(),
	)
	
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	task.ID = id
	return nil
}

//ユーザーIDからタスクを探す
func (r taskRepository) FindByUserID(userID int64) ([]*model.Task, error) {
	var taskList []*model.Task
	rows, err := r.db.Query(`SELECT id, user_id, title, description, priority, due_date, completed_at, created_at, updated_at FROM tasks WHERE user_id = ? ORDER BY created_at DESC`, userID)

	if err != nil {
    	return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		task := &model.Task{}
		rows.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.Description,
			&task.Priority,
			&task.DueDate,
			&task.CompletedAt,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		taskList = append(taskList, task)
	}
	return taskList, nil
}

//IDから1行取得
func (r taskRepository) FindByID(id int64, userID int64) (*model.Task, error) {
	task := &model.Task{}

	row := r.db.QueryRow(`SELECT id, user_id, title, description, priority, due_date, completed_at, created_at, updated_at FROM tasks WHERE id = ? AND user_id = ?`, id, userID)
	err := row.Scan(
		&task.ID,
		&task.UserID,
		&task.Title,
		&task.Description,
		&task.Priority,
		&task.DueDate,
		&task.CompletedAt,
		&task.CreatedAt,
		&task.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return task, nil
}

//タスクの更新
func (r taskRepository) Update(task *model.Task) error {
	_, err := r.db.Exec(`UPDATE tasks SET title = ?, description = ?, priority = ?, due_date = ?, completed_at = ?, updated_at = ? WHERE id = ? AND user_id = ?`, 
		task.Title,
		task.Description,
		task.Priority,
		task.DueDate,
		task.CompletedAt,
		time.Now(),
		task.ID, 
		task.UserID,
	)

	if err != nil {
		return err
	}
	return nil
}

//タスクの削除
func (r taskRepository) Delete(id int64, userID int64) error {
	_, err := r.db.Exec(`DELETE FROM tasks WHERE id = ? AND user_id = ?`, id, userID)

	if err != nil {
		return err
	}
	return nil
}