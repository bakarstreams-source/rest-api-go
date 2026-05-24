# REST API with Go and PostgreSQL

A RESTful API built with Go, PostgreSQL and JWT Authentication.

## Features
- User Registration with password hashing (bcrypt)
- User Login with JWT Authentication
- Protected routes with middleware
- Full CRUD operations on users

## Tech Stack
- Go
- PostgreSQL
- JWT (golang-jwt)
- bcrypt
- net/http

## API Endpoints

### Auth (Public)
- POST /register - Register new user
- POST /login - Login and get JWT token

### Users (Protected - JWT required)
- GET /users - Get all users
- GET /users/name - Get user by name
- POST /users - Create users
- PUT /users/id - Update user
- DELETE /users/id - Delete user
- DELETE /users - Delete all users

## Setup
1. Clone the repo
2. Install PostgreSQL
3. Create .env file:
DB_USER=your_username
DB_PASSWORD=your_password
DB_HOST=localhost
DB_PORT=5432
DB_NAME=restapi
JWT_SECRET=your_secret_key

4. Create database and table:
CREATE DATABASE restapi;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    password VARCHAR(255)
);

5. Run:
go run main.go

## Authentication
All /users routes require JWT token in header:
Authorization: Bearer your_token