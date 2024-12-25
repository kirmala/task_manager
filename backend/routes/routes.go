package routes

import (
	"net/http"
	"task-manager/controllers"
	"task-manager/middlewares"

	"github.com/gorilla/mux"
)

// Set up the router

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.Handle("/tasks", middlewares.JWTMiddleware(http.HandlerFunc(controllers.CreateTask))).Methods("POST")
	router.Handle("/tasks", middlewares.JWTMiddleware(http.HandlerFunc(controllers.GetTask))).Methods("GET")
	router.Handle("/tasks/{id}", middlewares.JWTMiddleware(http.HandlerFunc(controllers.UpdateTask))).Methods("PUT")
	router.Handle("/tasks/{id}", middlewares.JWTMiddleware(http.HandlerFunc(controllers.DeleteTask))).Methods("DELETE")

	
	return router
}
