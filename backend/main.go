package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yrk9/cardboard_task_app/db"
	"github.com/yrk9/cardboard_task_app/handler"
	"github.com/yrk9/cardboard_task_app/repository"
	"github.com/yrk9/cardboard_task_app/service"
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

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/v1/auth/login", authHandler.Login)

	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}