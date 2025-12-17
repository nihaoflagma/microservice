package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go-microservice/handlers"
	"go-microservice/metrics"
	"go-microservice/services"
	"go-microservice/utils"
)

func main() {
	metrics.Init()

	service := services.NewUserService()
	handler := &handlers.UserHandler{Service: service}

	r := mux.NewRouter()
	r.Use(utils.RateLimitMiddleware)
	r.Use(metrics.Middleware)

	r.HandleFunc("/api/users", handler.GetAll).Methods("GET")
	r.HandleFunc("/api/users/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/api/users", handler.Create).Methods("POST")
	r.HandleFunc("/api/users/{id}", handler.Update).Methods("PUT")
	r.HandleFunc("/api/users/{id}", handler.Delete).Methods("DELETE")

	r.Handle("/metrics", metrics.Handler())

	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
