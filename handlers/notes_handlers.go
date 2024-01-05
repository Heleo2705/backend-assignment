package handlers

import (
	"encoding/json"
	"fmt"
	// "log"
	"net/http"

	db "example.com/backend-assignment/db/sqlc"
	// "google.golang.org/genproto/googleapis/cloud/aiplatform/v1/schema/predict/params"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	// userId := r.Context().Value("userID")
	params := db.CreateNoteParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, "Error during parsing params")
		return
	}
	note, err := dbQueries.CreateNote(r.Context(), params)
	if err != nil {
		respondWithErr(w, 400, "Error during creating Note")
		return
	}
	respondWithJson(w, 200, note)
}

func GetNotesForUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID")
	id:= userId.(int)
	fmt.Printf("%v",id)
	// if ok == false {
	// 	log.Fatal(userId)
	// 	respondWithErr(w, 400, "Can't determine user id from token")
	// 	return
	// }
	notes, err := dbQueries.GetNotesForUser(r.Context(),int64(id))
	if err != nil {
		respondWithErr(w, 400, "Error getting notes for user")
		return
	}
	respondWithJson(w, 200, notes)
}
