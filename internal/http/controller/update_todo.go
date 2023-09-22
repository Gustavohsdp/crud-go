package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gustavohsdp/fo-api-postgresql/internal/domain/entity"
	"github.com/Gustavohsdp/fo-api-postgresql/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func UpdateTodoController(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Error ao fazer o parse do id: %v", err)

		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	var todo entity.Todo

	err = json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error ao fazer o decode do json: %v", err)

		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	rows, err := usecase.UpdateTodoUseCase(int64(id), todo)

	if err != nil {
		log.Printf("Error ao atualizar todo: %v", err)

		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if rows > 1 {
		log.Printf("Error: foram atualizados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Dados atualizados com sucesso.",
	}

	response.Header().Add("Content-type", "application/json")
	json.NewEncoder(response).Encode(resp)
}
