package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	db "example.com/backend-assignment/db/sqlc"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/meilisearch/meilisearch-go"
)

var dbQueries *db.Queries
var meiliConfig MeiliConfig

func InitDB() error {

	var err error
	err = godotenv.Load()
	if err != nil {
		return errors.New(err.Error())
	}
	dbDriver := os.Getenv("DB_DRIVER")
	dbSource := os.Getenv("DB_SOURCE")
	if dbDriver == "" {
		return errors.New("db Driver not found")
	}
	if dbSource == "" {
		return errors.New("db source not found")
	}
	dbCon, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to DB:", err)
		return err
	}
	fmt.Println("Initialising DB")
	dbQueries = db.New(dbCon)
	// dbStore = db.NewStore(dbCon)
	migrationUrl := os.Getenv("MIGRATION_URL")
	if migrationUrl == "" {
		return errors.New("migration url not found")
	}
	// runDBMigration(migrationUrl,dbSource)
	return nil
}

func InitMeili() error {
	godotenv.Load()
	fmt.Println("Starting Meili")
	meili_host := os.Getenv("MEILI_HOST")
	if meili_host == "" {
		return fmt.Errorf("can't find meili host")
	}
	meili_key := os.Getenv("MEILI_KEY")
	if meili_key == "" {
		return fmt.Errorf("can't find meili key")
	}
	meiliClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   meili_host,
		APIKey: meili_key,
	})
	var err error
	_, err = meiliClient.CreateIndex(&meilisearch.IndexConfig{
		Uid:        "notes",
		PrimaryKey: "id",
	})
	if err != nil {
		if strings.Contains(err.Error(), "index already exists") {
			fmt.Println("notes index already exists, continuing")
			return nil
		}
		if strings.Contains(err.Error(), "connection refused") {
			return fmt.Errorf("meili connection refused %v", err)
		}
		return err
	}
	meiliNoteIndex := meiliClient.Index("notes")
	var settingErr error
	_, settingErr = meiliNoteIndex.UpdateFilterableAttributes(
		&[]string{"id", "user_id"},
	)
	if settingErr != nil {
		return fmt.Errorf("error while updating filter settings %v", settingErr)
	}
	_, settingErr = meiliNoteIndex.UpdateSearchableAttributes(
		&[]string{"content"},
	)
	if settingErr != nil {
		return fmt.Errorf("error while updating search settings %v", settingErr)
	}
	meiliConfig = MeiliConfig{
		Client:    meiliClient,
		NoteIndex: meiliNoteIndex,
	}
	return nil
}
