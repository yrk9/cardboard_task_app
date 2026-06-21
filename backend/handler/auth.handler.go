package handler

import (
	"encoding/json"
	"net/http"
	"github.com/yrk9/cardboard_task_app/service"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

//リクエストボディ
type authRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

//レスポンスボディ
type authResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	//リクエストボディをauthRequestにデコード
	var req authRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//登録メソッドを呼ぶ
	token, err := h.authService.Register(req.Email, req.Password)
	//エラー処理
	if err == service.ErrUserAlreadyExists {
		w.WriteHeader(http.StatusConflict)
		return
	} else if (err == service.ErrEmailType) || (err == service.ErrPasswordLength) || (err == service.ErrEmailNil) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//成功
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authResponse{Token: token})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	//リクエストボディをauthRequestにデコード
	var req authRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := h.authService.Login(req.Email, req.Password)
	//エラー処理
	if err == service.ErrInvalidCredentials {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//成功
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(authResponse{Token: token})
}