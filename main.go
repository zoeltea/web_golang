package main

import (
	"account-service/database"
	"account-service/services/be"
	"account-service/services/fe"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Load .env file
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Parse command line arguments
	flag.Parse()

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set up routes BE
	be.AccountRoutes(db)
	fe.HomeRoutes()
	fe.AccountRoutes()

	// port := os.Getenv("SERVER_PORT")
	port := "8080"
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
