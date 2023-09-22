package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Gustavohsdp/fo-api-postgresql/internal/usecase"
	"github.com/go-chi/chi/v5"
)

func FindByIdTodoController(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Error ao fazer o parse do id: %v", err)

		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	todo, err := usecase.FindByIdTodoUseCase(int64(id))

	if err != nil {
		log.Printf("Error ao obter todo: %v", err)

		http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	response.Header().Add("Content-type", "application/json")
	json.NewEncoder(response).Encode(todo)
}
