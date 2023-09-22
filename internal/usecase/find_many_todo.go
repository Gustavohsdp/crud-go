package usecase

import (
	"github.com/Gustavohsdp/fo-api-postgresql/internal/domain/entity"
	"github.com/Gustavohsdp/fo-api-postgresql/internal/infra/database"
)

func FindManyTodoUseCase() (todos []entity.Todo, err error) {
	connection, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer connection.Close()

	sql := `SELECT * FROM todos`

	rows, err := connection.Query(sql)

	if err != nil {
		return
	}

	for rows.Next() {
		var todo entity.Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
