package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/db"
	"rest-api/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// DB se user dhundho email se
	var user models.User
	row := db.DB.QueryRow("SELECT id, name, email FROM users WHERE email = $1", req.Email)
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Token banao
	token, err := GenerateJWT(user.Email)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}