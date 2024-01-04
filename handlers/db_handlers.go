package handlers

import (
	"database/sql"
	"errors"
	db "example.com/backend-assignment/db/sqlc"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var dbQueries *db.Queries

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
