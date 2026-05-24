package models

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}


type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
