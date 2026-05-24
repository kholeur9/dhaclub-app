package main

import (
	//"fmt"
	//"html"
	//"context"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kholeur9/dhaclub-app/internal/db"
	"github.com/kholeur9/dhaclub-app/internal/helpers"
	"github.com/kholeur9/dhaclub-app/internal/todo"
	"github.com/kholeur9/dhaclub-app/internal/user"
)

func main() {
	godotenv.Load()
	port := ":8080"
	pg := db.Connect()

	router := http.NewServeMux()

	TodoStore := todo.NewPostgresTodo(pg)
	TodoService := todo.NewTodoService(TodoStore)
	HandlerTodo := todo.NewHandlerTodo(TodoService)

	router.HandleFunc("POST /todo", HandlerTodo.CreateTodoHandler)
	router.HandleFunc("GET /todos/{id}", HandlerTodo.GetTodoByIDHandler)
	router.HandleFunc("GET /todos", HandlerTodo.TodosListHandler)
	router.HandleFunc("DELETE /todos/{id}", HandlerTodo.DeleteTodoHandler)
	router.HandleFunc("PATCH /todos/{id}", HandlerTodo.UpdateTodoHandler)

	//router.Handle("GET /", http.FileServer(http.Dir("static")))
	UserStore := user.NewPostgresUser(pg)
	Secure := helpers.NewBcryptSecure()
	UserService := user.NewUserService(UserStore, Secure)
	UserHandler := user.NewUserHandler(UserService)
	router.HandleFunc("POST /user", UserHandler.CreateUserHandler)
	router.HandleFunc("PATCH /user/{id}", UserHandler.EmailUpdateHandler)

	log.Println("Starting server on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
