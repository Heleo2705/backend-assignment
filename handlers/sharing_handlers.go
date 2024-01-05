package handlers

import (
	// "encoding/json"
	// "encoding/json"
	"encoding/json"
	"net/http"
	"strconv"

	db "example.com/backend-assignment/db/sqlc"
	"github.com/go-chi/chi/v5"
)

func ShareNote(w http.ResponseWriter, r *http.Request) {
	type shareNotesParams struct {
		SharedNoteID int64 `json:"shared_note_id"`
	}
	userId := r.Context().Value("userID")
	id := userId.(int)
	otherIdParam := chi.URLParam(r, "id")
	sharedId, err := strconv.Atoi(otherIdParam)
	if err != nil {
		respondWithErr(w, 400, "Id is not integer")
		return
	}
	decoder := json.NewDecoder(r.Body)
	params := shareNotesParams{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, "Error parsing params")
		return
	}
	note, err := dbQueries.ShareNote(r.Context(), db.ShareNoteParams{OwnerID: int64(id), SharedID: int64(sharedId), SharedNoteID: params.SharedNoteID})
	if err != nil {
		respondWithErr(w, 400, "Error sharing notes")
		return
	}
	respondWithJson(w, 200, note)
}

func GetSharedNotes(w http.ResponseWriter, r *http.Request) {
	// decoder:=json.NewDecoder(r.Body)
	userId := r.Context().Value("userID")
	id := userId.(int)
	notes, err := dbQueries.GetSharedNotes(r.Context(), int64(id))
	if err != nil {
		respondWithErr(w, 400, "Error getting shared Notes")
		return
	}
	respondWithJson(w, 200, notes)
}
