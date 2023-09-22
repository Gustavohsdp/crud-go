package database

import (
	"database/sql"
	"fmt"

	"github.com/Gustavohsdp/fo-api-postgresql/config"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	configuration := config.GetDB()

	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configuration.Host, configuration.Port, configuration.User, configuration.Pass, configuration.Database,
	)

	connection, err := sql.Open("postgres", stringConnection)

	if err != nil {
		return nil, err
	}

	err = connection.Ping()

	return connection, err
}
