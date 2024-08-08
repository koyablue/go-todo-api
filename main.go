package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:userpassword@tcp(go-todo-mysql:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	fmt.Println("Starting server on :8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!aaaaaaaaaaaaaaaaaaaaaaa")
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
