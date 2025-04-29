package be

import (
	"account-service/handlers"
	"database/sql"
	"net/http"
)

func AccountRoutes(db *sql.DB) {
	http.HandleFunc("/api/accounts", handlers.GetAccounts(db))
	http.HandleFunc("/api/accounts/register", handlers.RegisterAccount(db))
	// http.HandleFunc("/api/accounts/register", handlers.RegisterAccount(db))
	// http.HandleFunc("/api/accounts/update", handlers.UpdateAccount(db))
	// http.HandleFunc("/api/accounts/delete", handlers.DeleteAccount(db))
}
