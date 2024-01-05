package handlers

import (
	"encoding/json"
	"net/http"

	db "example.com/backend-assignment/db/sqlc"
	"example.com/backend-assignment/utils"
	// "google.golang.org/genproto/googleapis/cloud/aiplatform/v1/schema/predict/params"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := db.CreateUserParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, "Error parsing params")
		return
	}
	_, err2 := dbQueries.GetUser(r.Context(), params.Name)
	if err2 == nil {
		respondWithErr(w, 400, "User Already exists")
		return
	}
	user, err := dbQueries.CreateUser(r.Context(), params)
	if err != nil {
		respondWithErr(w, 400, err.Error())
		return
	}
	respondWithJson(w, 200, user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	params := db.CreateUserParams{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, 400, "Error parsing params")
		return
	}
	user, err := dbQueries.GetUser(r.Context(), params.Name)
	if err != nil {
		respondWithErr(w, 400, err.Error())
		return
	}
	if params.Password != user.Password {
		respondWithErr(w, 404, "Invalid Password")
		return
	}

	token, err := utils.GenerateToken(int(user.ID))
	type outputToken struct {
		Token string `json:"name"`
	}
	respondWithJson(w, 200, outputToken{Token: token})

}
