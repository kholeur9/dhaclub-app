package main

import (
	//"fmt"
	//"html"
	//"context"
	"log"
	"net/http"

	"github.com/kholeur9/dhaclub-app/internal/db"
	"github.com/kholeur9/dhaclub-app/internal/todo"
)

func main() {
	port := ":8080"

	pg := db.Connect()

	router := http.NewServeMux()

	store := todo.NewPostgresTodo(pg)
	TodoService := todo.NewTodoService(store)
	HandlerTodo := todo.NewHandlerTodo(TodoService)

	router.HandleFunc("POST /todo", HandlerTodo.CreateTodoHandler)
	router.HandleFunc("GET /todo/{todo_id}", HandlerTodo.GetTodoByIDHandler)
	router.HandleFunc("GET /todos", HandlerTodo.TodosListHandler)

	//router.Handle("GET /", http.FileServer(http.Dir("static")))

	log.Println("Starting server on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
