package usecase

import (
	"github.com/Gustavohsdp/fo-api-postgresql/internal/infra/database"
)

func DeleteTodoUseCase(id int64) (int64, error) {
	connection, err := database.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer connection.Close()

	sql := `DELETE FROM todos WHERE id=$1`

	res, err := connection.Exec(sql, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
