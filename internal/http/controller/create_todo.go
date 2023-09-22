package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Gustavohsdp/fo-api-postgresql/internal/domain/entity"
	"github.com/Gustavohsdp/fo-api-postgresql/internal/usecase"
)

func CreateTodoController(response http.ResponseWriter, request *http.Request) {
	var todo entity.Todo

	err := json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error ao fazer o decode do json: %v", err)

		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	id, err := usecase.CreateTodoUseCase(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar criar: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo criado com sucesso! ID: %d", id),
		}
	}
	response.Header().Add("Content-type", "application/json")
	json.NewEncoder(response).Encode(resp)
}
