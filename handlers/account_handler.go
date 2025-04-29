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

		rows, err := db.Query("SELECT id, name, nik, no_hp, no_rekening, saldo FROM accounts")
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
