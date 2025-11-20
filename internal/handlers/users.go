package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"uber-app-backend/internal/db"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	ClerkID string `json:"clerkId"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if u.Name == "" || u.Email == "" || u.ClerkID == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	_, err := db.Pool.Exec(
		context.Background(),
		`INSERT INTO users (name, email, clerk_id) VALUES ($1, $2, $3)`,
		u.Name, u.Email, u.ClerkID,
	)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
