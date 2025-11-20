package main

import (
	"log"
	"net/http"

	"uber-app-backend/internal/db"
	"uber-app-backend/internal/handlers"
)

func main() {
	db.Init()

	http.HandleFunc("/api/user", handlers.CreateUser)
	http.HandleFunc("/api/drivers", handlers.GetDrivers)
	http.HandleFunc("/api/ride/create", handlers.CreateRide)

	log.Println("🚀 Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
