package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Msg string `json:"message"`
}

func Response(w http.ResponseWriter, status int, response any) {
	bytes, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(bytes)
}
