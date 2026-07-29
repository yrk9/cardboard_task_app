package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yrk9/cardboard_task_app/db"
	"github.com/yrk9/cardboard_task_app/handler"
	"github.com/yrk9/cardboard_task_app/repository"
	"github.com/yrk9/cardboard_task_app/service"
	"github.com/yrk9/cardboard_task_app/middleware"
)

func main() {
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatal("DB接続失敗: ", err)
	}
	defer dbConn.Close()

	userRepo := repository.NewUserRepository(dbConn)
	authService := service.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	taskRepo := repository.NewTaskRepository(dbConn)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/v1/auth/login", authHandler.Login)

	mux.Handle("GET /api/v1/tasks", middleware.AuthMiddleware(http.HandlerFunc(taskHandler.List)))
	mux.Handle("POST /api/v1/tasks", middleware.AuthMiddleware(http.HandlerFunc(taskHandler.Create)))
	mux.Handle("GET /api/v1/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(taskHandler.Get)))
	mux.Handle("PUT /api/v1/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(taskHandler.Update)))
	mux.Handle("DELETE /api/v1/tasks/{id}", middleware.AuthMiddleware(http.HandlerFunc(taskHandler.Delete)))
	mux.Handle("PATCH /api/v1/tasks/{id}/complete", middleware.AuthMiddleware(http.HandlerFunc(taskHandler.ToggleComplete)))

	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}