package usecase

import (
	"github.com/Gustavohsdp/fo-api-postgresql/internal/domain/entity"
	"github.com/Gustavohsdp/fo-api-postgresql/internal/infra/database"
)

func FindByIdTodoUseCase(id int64) (todo entity.Todo, err error) {
	connection, err := database.OpenConnection()

	if err != nil {
		return
	}

	defer connection.Close()

	sql := `SELECT * FROM todos WHERE id=$1`

	err = connection.QueryRow(sql, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done)

	return
}
