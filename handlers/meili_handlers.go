package handlers

import (
	// "fmt"

	"fmt"

	db "example.com/backend-assignment/db/sqlc"
	"github.com/meilisearch/meilisearch-go"
	// "github.com/sqlc-dev/pqtype"
)

type MeiliConfig struct {
	Client    *meilisearch.Client
	NoteIndex *meilisearch.Index
}

func (config MeiliConfig) AddNoteToMeili(note db.Note) error {

	_, err := config.NoteIndex.AddDocuments([]db.Note{note})
	if err != nil {
		return fmt.Errorf("error while adding note to meili %v %v", err, note)
	}
	return nil
}

func (config MeiliConfig) UpdateNoteInMeili(note db.Note) error {

	_, err := config.NoteIndex.UpdateDocuments([]db.Note{note})
	if err != nil {
		return fmt.Errorf("error while updating note to meili %v %v", err, note)
	}
	return nil
}
func (config MeiliConfig) DeletingNoteinMeili(note db.Note) error {
	filterStr := fmt.Sprintf("%v", note.ID)
	_, err := config.NoteIndex.DeleteDocument(filterStr)
	if err != nil {
		return fmt.Errorf("error while updating note to meili %v %v", err, note)
	}
	return nil
}

func (config MeiliConfig) SearchNote(user_id int64, query string) ([]interface{}, error) {

	filterStr := fmt.Sprintf("user_id=%v", user_id)
	fmt.Println(query)
	// return nil,nil
	resp, err := config.NoteIndex.Search(query, &meilisearch.SearchRequest{
		Filter: []string{filterStr},
	
	})
	if err != nil {
		return nil, err
	}
	hits := resp.Hits
	return hits, nil
}
