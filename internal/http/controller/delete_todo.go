package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gustavohsdp/fo-api-postgresql/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func DeleteTodoController(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Error ao fazer o parse do id: %v", err)

		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	rows, err := usecase.DeleteTodoUseCase(int64(id))

	if err != nil {
		log.Printf("Error ao deletar todo: %v", err)

		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	if rows > 1 {
		log.Printf("Error: foram removidos %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Todo removido com sucesso.",
	}

	response.Header().Add("Content-type", "application/json")
	json.NewEncoder(response).Encode(resp)
}
