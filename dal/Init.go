package dal

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const (
	forumDbPath    = "./dal/data/forum.db"
	identityDbPath = "./dal/data/identity.db"
	dbDriver       = "sqlite3"
)

var (
	forumDb    *sql.DB
	identityDb *sql.DB
)

func Init() {
	log.Println("dal: initialising...")
	var err error
	log.Println("dal: opening identity database...")
	identityDb, err = sql.Open(dbDriver, identityDbPath)
	if err != nil {
		log.Fatalf("dal: error opening identity database:\n%s\n", err.Error())
	}
	log.Println("dal: creating users table...")
	err = createUsersTable(identityDb)
	if err != nil {
		log.Fatalf("dal: error creating users table:\n%s\n", err.Error())
	}
	log.Println("dal: opening forum database...")
	forumDb, err = sql.Open(dbDriver, forumDbPath)
	if err != nil {
		log.Fatalf("dal: error opening forum database:\n%s\n", err.Error())
	}
	log.Println("dal: initialised successfully.")
}

func createUsersTable(db *sql.DB) error {
	queryStr := `CREATE TABLE IF NOT EXISTS "users" (
		"id" INTEGER PRIMARY KEY AUTOINCREMENT,
		"nickname" TEXT NOT NULL UNIQUE,
		"age" TEXT NOT NULL,
		"gender" TEXT NOT NULL,
		"firstname" TEXT NOT NULL,
		"lastname" TEXT NOT NULL,
		"email" TEXT NOT NULL UNIQUE,
		"encryptedPassword" TEXT NOT NULL);`
	statement, err := db.Prepare(queryStr)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}
