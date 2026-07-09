package main

import (
	//"fmt"
	//"html"
	//"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	auth_handler "github.com/kholeur9/dhaclub-app/internal/auth/handler"
	"github.com/kholeur9/dhaclub-app/internal/auth/jwt"
	auth_service "github.com/kholeur9/dhaclub-app/internal/auth/service"
	"github.com/kholeur9/dhaclub-app/internal/db"
	"github.com/kholeur9/dhaclub-app/internal/helpers"
	"github.com/kholeur9/dhaclub-app/internal/todo"
	"github.com/kholeur9/dhaclub-app/internal/user"
)

func main() {
	godotenv.Load()
	port := ":8080"
	pg := db.Connect()
	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		log.Fatal("JWT_SECRET is missing")
	}

	router := http.NewServeMux()

	//router.Handle("GET /", http.FileServer(http.Dir("static")))
	UserStore := user.NewPostgresUser(pg)
	Secure := helpers.NewBcryptSecure()
	UserService := user.NewUserService(UserStore, Secure)
	UserHandler := user.NewUserHandler(UserService)
	router.HandleFunc("POST /register", UserHandler.CreateUserHandler)

	//JWT
	JwtService := jwt.NewJwtService(jwtKey)
	Middleware := jwt.NewAuthMiddlewareService(JwtService)
	AuthService := auth_service.NewAuthService(UserService, JwtService)
	AuthHandler := auth_handler.NewAuthHandler(AuthService)
	router.HandleFunc("POST /login", AuthHandler.LoginUserHandler)

	router.Handle("GET /me", Middleware.AuthMiddleware(http.HandlerFunc(UserHandler.GetMeHandler)))
	router.Handle("PATCH /user/{id}", Middleware.AuthMiddleware(http.HandlerFunc(UserHandler.UserUpdateHandler)))
	router.Handle("PATCH /user/{id}/change-password", Middleware.AuthMiddleware(http.HandlerFunc(UserHandler.UpdatePasswordHandler)))

	// Todo User
	TodoStore := todo.NewPostgresTodo(pg)
	TodoService := todo.NewTodoService(TodoStore)
	HandlerTodo := todo.NewHandlerTodo(TodoService)
	router.Handle("POST /todo", Middleware.AuthMiddleware(http.HandlerFunc(HandlerTodo.CreateTodoHandler)))
	router.Handle("GET /todos/{id}", Middleware.AuthMiddleware(http.HandlerFunc(HandlerTodo.GetTodoByIDHandler)))
	router.Handle("GET /todos", Middleware.AuthMiddleware(http.HandlerFunc(HandlerTodo.TodosListHandler)))	
	router.Handle("DELETE /todos/{id}", Middleware.AuthMiddleware(http.HandlerFunc(HandlerTodo.DeleteTodoHandler)))
	router.Handle("PATCH /todos/{id}", Middleware.AuthMiddleware(http.HandlerFunc(HandlerTodo.UpdateTodoHandler)))

	log.Println("Starting server on port", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
