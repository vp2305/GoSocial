package main

import (
	"SocialMedia/internal/models"
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) handleCreateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}

	// Validate inputs (basic validation, can be expanded)
	if input.Username == "" || input.Email == "" || input.Password == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	if len(input.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	newUser := &models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password, // Make sure to hash the password before saving it
	}

	// Create the user in the store
	err = app.store.Users.Create(r.Context(), newUser)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	// Send a success response
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"message": "User created successfully",
		"data": map[string]interface{}{
			"id":         newUser.ID,
			"username":   newUser.Username,
			"email":      newUser.Email,
			"password":   newUser.Password,
			"created_at": newUser.CreatedAt,
		},
	}
	json.NewEncoder(w).Encode(response)
}
