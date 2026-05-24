package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/db"
	"rest-api/models"
	"golang.org/x/crypto/bcrypt"	

)





func Register(w http.ResponseWriter, r *http.Request) {
    var req models.User
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Password hash karo
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    // DB mein save karo
    var user models.User
    err = db.DB.QueryRow(
        "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id, name, email",
        req.Name, req.Email, string(hashedPassword),
    ).Scan(&user.ID, &user.Name, &user.Email)

    if err != nil {
        http.Error(w, "Error creating user", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var req models.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // DB se user dhundho
    var user models.User
    row := db.DB.QueryRow(
        "SELECT id, name, email, password FROM users WHERE email = $1",
        req.Email,
    )
    if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    // Password verify karo
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
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