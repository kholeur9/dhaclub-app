package main

import (
	//"fmt"
	//"html"
	//"context"
	"log"
	"net/http"

	"dhaclub-app/internal/todo"
)

func main() {
	port := ":8080"

	router := http.NewServeMux()

	store := todo.NewMemoryTodo()
	TodoService := todo.NewTodoService(store)
	HandlerTodo := todo.NewHandlerTodo(TodoService)

	router.HandleFunc("POST /create-todo", HandlerTodo.CreateTodoHandler)
	router.HandleFunc("GET /todo", HandlerTodo.GetTodoByIDHandler)
	router.HandleFunc("GET /todos", HandlerTodo.TodosListHandler)

	//router.Handle("GET /", http.FileServer(http.Dir("static")))

	log.Println("Starting server on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
