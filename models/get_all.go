package models

import "github.com/santosant/go-api/db"

func getAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query("SELECT id, title, description, completed FROM todos")
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return
}
