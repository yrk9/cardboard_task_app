package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"github.com/yrk9/cardboard_task_app/db"
)

func main() {
	dbConn, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	port := os.Getenv("SERVER_PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "200 OK")
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}