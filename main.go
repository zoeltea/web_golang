package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hallo dunia manca negara!")
}

func main() {
	port := "8080"
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handler)

	fmt.Println("Server berjalan...")
	fmt.Printf("Server mendengarkan pada port %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
