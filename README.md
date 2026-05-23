# REST API with Go and PostgreSQL

A RESTful API built with Go and PostgreSQL.

## Features
- Get all users
- Get user by ID
- Create users
- Update user
- Delete user

## Tech Stack
- Go
- PostgreSQL
- net/http

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | /users | Get all users |
| GET | /users/{id} | Get user by ID |
| POST | /users | Create users |
| PUT | /users/{id} | Update user |
| DELETE | /users/{id} | Delete user |

## Setup

1. Clone the repo
2. Create .env file:
DB_USER=your_username
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=restapi

3. Run:
go run main.go