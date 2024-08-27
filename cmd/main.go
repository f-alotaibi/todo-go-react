package main

import (
	"fmt"
	"net/http"
	"os"
	"todo_go_react/internal/handlers"
	"todo_go_react/internal/storage"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "dist/index.html")
	})
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("dist/assets"))))

	store := storage.NewMemoryStorage()
	/* Test values
	store.Create(models.TodoTask{
		Text: "First",
	})
	store.Create(models.TodoTask{
		Text: "Second",
	})
	*/

	todoHandler := handlers.NewTodoHandler(store)

	http.HandleFunc("GET /tasks", todoHandler.GetTodos)
	http.HandleFunc("POST /tasks", todoHandler.AddTodo)
	http.HandleFunc("PUT /tasks/{id}", todoHandler.UpdateTodo)
	http.HandleFunc("DELETE /tasks/{id}", todoHandler.DeleteTodo)

	port := os.Getenv("PORT")
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
