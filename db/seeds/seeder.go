package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:userpassword@tcp(go-todo-mysql:3306)/todo_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tasks := []string{
		"Task 1",
		"Task 2",
		"Task 3",
	}

	valueStrings := make([]string, 0, len(tasks))
	valueArgs := make([]interface{}, 0, len(tasks))

	for _, task := range tasks {
		valueStrings = append(valueStrings, "(?)")
		valueArgs = append(valueArgs, task)
	}

	stmt := fmt.Sprintf("Insert INTO todos (task) VALUES %s", strings.Join(valueStrings, ","))
	_, err = db.Exec(stmt, valueArgs...)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Seeder executed successfully")
}
