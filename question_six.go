package main

import (
	"fmt"
	"net/http"
)

// HelloServer handles incoming HTTP requests and writes "Hello, World!" as the response.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func main() {
	// Registering HelloServer as the handler function for the root URL path ("/").
	http.HandleFunc("/", HelloServer)

	// Printing a message indicating that the server is listening on port 8080.
	fmt.Println("Server listening on port 8080...")

	// Starting the HTTP server on port 8080.
	err := http.ListenAndServe(":8080", nil)

	// Handling any error that occurs while starting the server.
	if err != nil {
		panic(err)
	}
}
