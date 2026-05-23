package db

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func OpenDB() (*sql.DB, error) {
    godotenv.Load()

    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_NAME"),
    )

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    if err = db.Ping(); err != nil {
        return nil, err
    }

    DB = db
    return db, nil
}

func CloseDB() {
    if DB != nil {
        DB.Close()
    }
}