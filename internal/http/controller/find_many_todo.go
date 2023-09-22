package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Gustavohsdp/fo-api-postgresql/internal/usecase"
)

func FindManyTodoController(response http.ResponseWriter, request *http.Request) {

	todos, err := usecase.FindManyTodoUseCase()

	if err != nil {
		log.Printf("Erro ao obter todos: %v", err)
	}

	response.Header().Add("Content-type", "application/json")
	json.NewEncoder(response).Encode(todos)
}
