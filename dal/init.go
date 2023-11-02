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

func init() {
	log.Println("dal: initialising...")
	var err error
	log.Println("dal: opening identity database...")
	identityDb, err = sql.Open(dbDriver, identityDbPath)
	if err != nil {
		log.Println("dal: error opening identity database:")
		log.Fatal(err.Error())
	}
	log.Println("dal: opening forum database...")
	forumDb, err = sql.Open(dbDriver, forumDbPath)
	if err != nil {
		log.Println("dal: error opening forum database:")
		log.Fatal(err.Error())
	}
	log.Println("dal: initialised successfully.")
}
