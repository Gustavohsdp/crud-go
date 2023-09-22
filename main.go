package main

import (
	"fmt"
	"net/http"

	"github.com/Gustavohsdp/fo-api-postgresql/config"
	"github.com/Gustavohsdp/fo-api-postgresql/internal/http/controller"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := config.Load()

	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/todo", controller.CreateTodoController)
	router.Put("/todo/{id}", controller.UpdateTodoController)
	router.Get("/todo", controller.FindManyTodoController)
	router.Get("/todo/{id}", controller.FindByIdTodoController)
	router.Delete("/todo/{id}", controller.DeleteTodoController)

	http.ListenAndServe(fmt.Sprintf(":%s", config.GetServerPort()), router)
}
