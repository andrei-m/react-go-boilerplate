package boilerplate

import (
	"log"
	"net/http"
)

func StartServer() error {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("listening on port 8080")
	return http.ListenAndServe(":8080", nil)
}
