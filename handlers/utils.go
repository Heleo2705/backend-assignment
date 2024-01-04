package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Msg string `json:"error_msg"`
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json Reponse %v", payload)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	w.Write(dat)
}

func respondWithErr(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error", msg)
	}
	respondWithJson(w, code, Error{Msg: msg})
}
