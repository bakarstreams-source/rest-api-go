package main

import (
	"net/http"
	"rest-api/db"
	"rest-api/handlers"
	"log"
)
func main () {
	  // DB pehle connect karo
    _, err := db.OpenDB()
    if err != nil {
        log.Fatal("DB connect nahi hua:", err)
    }
	defer db.CloseDB() // server band hone pe automatically close hoga

	// Mux create karo
	mux := http.NewServeMux()
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        handlers.CreateUser(w, r)
    }else if r.Method == http.MethodDelete {
		handlers.DeleteAllUsers(w, r)
	} else {
        handlers.GetAllUsers(w, r)
    }

})

mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodDelete {
        handlers.DeleteUserByID(w, r)
    } else if r.Method == http.MethodPut {
        handlers.UpdateUserByID(w, r)
    } else if r.Method == http.MethodGet {
        handlers.GetUserByName(w, r)
    }
})


	// Server start karo
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Server start nahi hua:", err)
	}

}
