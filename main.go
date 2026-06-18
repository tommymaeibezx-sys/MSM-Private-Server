package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback local
	}

	http.HandleFunc("/", handler)

	fmt.Println("Servidor corriendo en puerto:", port)

	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
