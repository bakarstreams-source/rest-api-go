# REST API with Go and PostgreSQL

A RESTful API built with Go, PostgreSQL and JWT Authentication.

## Features
- Get all users
- Get user by ID
- Create users
- Update user
- Delete user
- JWT Authentication

## Tech Stack
- Go
- PostgreSQL
- JWT
- net/http

## API Endpoints
- POST /login - Login and get token
- GET /users - Get all users (protected)
- GET /users/id - Get user by ID (protected)
- POST /users - Create users (protected)
- PUT /users/id - Update user (protected)
- DELETE /users/id - Delete user (protected)

## Setup
1. Clone the repo
2. Create .env file:
DB_USER=your_username
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=restapi
JWT_SECRET=your_secret_key

3. Run:
go run main.go