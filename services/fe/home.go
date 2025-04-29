package fe

import (
	"net/http"
)

func HomeRoutes() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
}
