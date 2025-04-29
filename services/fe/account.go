package fe

import (
	"net/http"
)

func AccountRoutes() {
	// Serve account pages
	// accountFs := http.FileServer(http.Dir("./static/account"))
	// http.Handle("/account/", http.StripPrefix("/account", accountFs))

	http.Handle("/account", http.FileServer(http.Dir("./static/account")))
}
