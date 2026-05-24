package models

type User struct {
    ID       int
    Name     string
    Email    string
    Password string
}

type LoginRequest struct {
    Email    string
    Password string
}