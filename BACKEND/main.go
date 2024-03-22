/*main.go: This is the main entry point of your Go application. It initializes the HTTP server and defines routes.*/

package main

import (
	"github.com/Bhavya022/SKYSERVE-ASSIGNMENT/BACKEND/handlers"
	"github.com/Bhavya022/SKYSERVE-ASSIGNMENT/BACKEND/storage"
	"net/http"
)

func main() {
	// Initialize file storage
	fileStorage := storage.NewFileStorage("uploads")

	// Define HTTP routes
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/upload", handlers.UploadHandler) // Example route for file upload

	// Start the server
	http.ListenAndServe(":8080", nil)
}
