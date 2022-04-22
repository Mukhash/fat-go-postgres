package db

import (
	"database/sql"
	"go-postgres/models"
)

func InsertTask(task *models.Task) error {
	query := `
	INSERT INTO tasks (title, body)
	VALUES ($1, $2)`

	_, err := conn.Exec(query, task.Title, task.Body)
	if err != nil {
		return err
	}
	return nil
}

func GetTasks() ([]models.Task, error) {
	query := `
	SELECT id, title, body FROM tasks
	ORDER BY id DESC`

	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Body)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func GetTask(id int) (*models.Task, error) {
	query := `
	SELECT id, title, body FROM tasks
	WHERE id = $1`

	row := conn.QueryRow(query, id)

	var task models.Task
	err := row.Scan(&task.ID, &task.Title, &task.Body)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func UpdateTask(task *models.Task) error {
	query := `
	UPDATE tasks
	SET title = $2, body = $3
	WHERE id = $1
	`
	_, err := conn.Exec(query, task.ID, task.Title, task.Body)
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(id int) error {
	query := `
	DELETE FROM tasks
	WHERE id = $1`

	_, err := conn.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
