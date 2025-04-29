package handlers

import (
	"account-service/models"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetAccounts(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("SELECT id, name, nik, no_hp, no_rekening, balance FROM accounts")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var accounts []models.Account
		for rows.Next() {
			var a models.Account
			if err := rows.Scan(&a.ID, &a.Name, &a.NIK, &a.NoHP, &a.NoRekening, &a.Balance); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			accounts = append(accounts, a)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(accounts)
	}
}

// RegisterAccount handler
func RegisterAccount(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var account models.Account
		if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validate required fields
		if account.Name == "" || account.NIK == "" || account.NoHP == "" || account.NoRekening == "" {
			http.Error(w, "All fields are required", http.StatusBadRequest)
			return
		}

		// Set default balance if not provided
		if account.Balance == 0 {
			account.Balance = 0.0
		}

		// Insert into database
		err := db.QueryRow(
			`INSERT INTO account (name, nik, no_hp, no_rekening, balance) 
			 VALUES ($1, $2, $3, $4, $5) 
			 RETURNING id`,
			account.Name, account.NIK, account.NoHP, account.NoRekening, account.Balance,
		).Scan(&account.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)
	}
}
