package handler

import (
	"encoding/json"
	"net/http"
	"github.com/yrk9/cardboard_task_app/middleware"
	"github.com/yrk9/cardboard_task_app/service"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) parseIDParam(w http.ResponseWriter, r *http.Request) (int64, bool) {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		http.Error(w, "idを取得できませんでした", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	userID, ok := middleware.GetUserID(r)

	if ! ok {
		http.Error(w, "認証が必要です", http.StatusUnauthorized)
		return
	}
	taskList, err := h.taskService.List(userID)

	if err != nil {
		http.Error(w, "サーバーエラー", http.StatusInternalServerError)
		return
	}
	w.Header().set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(taskList)

}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	// userID取得
	userID, ok := middleware.GetUserID(r)

	if ! ok {
		http.Error(w, "認証が必要です", http.StatusUnauthorized)
		return
	}
	taskList, err := h.taskService.List(userID)

	//ボディをデコード
	var input service.TaskInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "リクエストの形式が不正です", http.StatusBadRequest)
		return
	}

	// service呼び出し
	err := h.taskService.Create(userID, input)   // 今は error だけ返す形でしたね
	if err != nil {
		switch {
		case errors.Is(err, service.ErrTitleRequired),
			errors.Is(err, service.ErrInvalidPriority):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "サーバーエラーが発生しました", http.StatusInternalServerError)
		}
		return
	}

	// 成功
	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	// userID取得 NGなら401
	userID, ok := middleware.GetUserID(r) 

	if !ok {
		http.Error(w, "認証が必要です", http.StatusUnauthorized)
		return
	}
	// URLの{id}をint64に変換して取得 NGなら400
	id, ok := h.parseIDParam(w, r)   
	if !ok {
		return
	}
	// taskService.Get(userID, id) を呼ぶ
	task, err := h.taskService.Get(userID, id)
	// エラー処理
	//    - ErrTaskNotFound → 404
	//    - その他 → 500
	if errors.Is(err, service.ErrTaskNotFound) {
		http.Error(w, "見つからない", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "サーバー内部のエラー", http.StatusInternalServerError)
		return	
	}

	// 成功 → 200 + タスクをJSONで返す
	w.Header().set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	// userID取得 NGなら401
	userID, ok := middleware.GetUserID(r) 

	if !ok {
		http.Error(w, "認証が必要です", http.StatusUnauthorized)
		return
	}
	// URLの{id}を取得 → int64変換 NGなら400
	id, ok := h.parseIDParam(w, r)   
	if !ok {
		return
	}
	// リクエストボディを service.TaskInput にデコード NGなら400
	var input taskService.TaskInput

	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "リクエストを取得できませんでした", http.StatusBadRequest)
		return
	}
	// taskService.Update(userID, id, input) を呼ぶ
	task, err := h.taskService.Update(userID, id, input)
	// エラー処理
	//    - ErrTaskNotFound → 404
	//    - ErrTitleRequired / ErrInvalidPriority → 400
	//    - その他 → 500
	if errors.Is(err, service.ErrTaskNotFound) {
		http.Error(w, "見つからない", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "サーバー内部のエラー", http.StatusInternalServerError)
		return	
	}

	// 成功 → 200（更新後のタスクを返すならJSONエンコード）
	w.Header().set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	// userID取得 NGなら401
	userID, ok := middleware.GetUserID(r) 

	if !ok {
		http.Error(w, "認証が必要です", http.StatusUnauthorized)
		return
	}
	// URLの{id}を取得 → int64変換 NGなら400
	id, ok := h.parseIDParam(w, r)   
	if !ok {
		return
	}
	// taskService.Delete(userID, id) を呼ぶ
	err = h.taskService.Delete(userID, id)
	// エラー処理
	//    - ErrTaskNotFound → 404
	//    - その他 → 500
	if errors.Is(err, service.ErrTaskNotFound) {
		http.Error(w, "見つからない", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "サーバー内部のエラー", http.StatusInternalServerError)
		return	
	}

	// 成功 → 204 No Content（ボディなし、WriteHeaderのみ）
	w.Header().set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) ToggleComplete(w http.ResponseWriter, r *http.Request) {
	// userID取得 NGなら401
	userID, ok := middleware.GetUserID(r) 

	if !ok {
		http.Error(w, "認証が必要です", http.StatusUnauthorized)
		return
	}
	// URLの{id}を取得 → int64変換 NGなら400
	id, ok := h.parseIDParam(w, r)   
	if !ok {
		return
	}
	// taskService.ToggleComplete(userID, id) を呼ぶ
	task, err := h.taskService.ToggleComplete(userID, id)
	// エラー処理
	//    - ErrTaskNotFound → 404
	//    - その他 → 500
	if errors.Is(err, service.ErrTaskNotFound) {
		http.Error(w, "見つからない", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "サーバー内部のエラー", http.StatusInternalServerError)
		return	
	}

	// 成功 → 200 + 更新後のタスクをJSONで返す
	w.Header().set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}
