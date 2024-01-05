package handlers

import (
	"encoding/json"
	"fmt"
	"time"

	// "log"
	"net/http"

	db "example.com/backend-assignment/db/sqlc"
	// "github.com/go-chi/chi/v5"
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
	err = meiliConfig.AddNoteToMeili(note)
	if err != nil {
		respondWithErr(w, 400, "Error adding Note to meili")
		return
	}
	respondWithJson(w, 200, note)
}

func GetNotesForUser(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userID")
	id := userId.(int)
	fmt.Printf("%v", id)
	// if ok == false {
	// 	log.Fatal(userId)
	// 	respondWithErr(w, 400, "Can't determine user id from token")
	// 	return
	// }
	notes, err := dbQueries.GetNotesForUser(r.Context(), int64(id))
	if err != nil {
		respondWithErr(w, 400, "Error getting notes for user")
		return
	}
	respondWithJson(w, 200, notes)
}

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	params := db.UpdateNoteParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, "Error parsing params")
		return
	}
	params.LastUpdated = time.Now()
	note, err := dbQueries.UpdateNote(r.Context(), params)
	if err != nil {
		respondWithErr(w, 400, "Error updating notes")
		return
	}
	err = meiliConfig.UpdateNoteInMeili(note)
	if err != nil {
		respondWithErr(w, 400, "Error updating Note to meili")
		return
	}
	respondWithJson(w, 200, note)
}

func DeleteNotes(w http.ResponseWriter, r *http.Request) {
	params := db.DeleteNoteParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, "Error parsing params")
		return
	}
	note, err := dbQueries.DeleteNote(r.Context(), params)
	if err != nil {
		respondWithErr(w, 400, "Error deleting notes")
		return
	}
	err = meiliConfig.DeletingNoteinMeili(note)
	if err != nil {
		respondWithErr(w, 400, "Error deleting Note to meili")
		return
	}
	respondWithJson(w, 200, note)
}

func GetNote(w http.ResponseWriter, r *http.Request) {
	params := db.GetNoteParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, "Error parsing params")
		return
	}
	note, err := dbQueries.GetNote(r.Context(), params)
	if err != nil {
		respondWithErr(w, 400, "Error getting note")
		return
	}
	respondWithJson(w, 200, note)
}

func SearchNotes(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query().Get("q")
	
	userId := r.Context().Value("userID")
	id := userId.(int)
	notes, err := meiliConfig.SearchNote(int64(id), queryString)
	if err != nil {
		respondWithErr(w, 400, err.Error())
		return
	}
	respondWithJson(w, 200, notes)

}
