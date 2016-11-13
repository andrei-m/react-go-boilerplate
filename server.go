package boilerplate

import (
	"encoding/json"
	"log"
	"net/http"
)

type Hello struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)
	if err := e.Encode(Hello{"Hello, world!"}); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func StartServer() error {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/hello", helloHandler)

	log.Println("listening on port 8080")
	return http.ListenAndServe(":8080", nil)
}
