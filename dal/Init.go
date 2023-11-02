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
	log.Println("dal: creating posts table...")
	err = createPostsTable(forumDb)
	if err != nil {
		log.Fatalf("dal: error creating posts table:\n%s\n", err.Error())
	}
	log.Println("dal: creating comments table...")
	err = createCommentsTable(forumDb)
	if err != nil {
		log.Fatalf("dal: error creating comments table:\n%s\n", err.Error())
	}
	log.Println("dal: initialised successfully.")
}

func createUsersTable(db *sql.DB) error {
	queryStr := `CREATE TABLE IF NOT EXISTS "users" (
		"ID" INTEGER PRIMARY KEY AUTOINCREMENT,
		"Nickname" TEXT NOT NULL UNIQUE,
		"Age" TEXT NOT NULL,
		"Gender" TEXT NOT NULL,
		"FirstName" TEXT NOT NULL,
		"LastName" TEXT NOT NULL,
		"EmailAddress" TEXT NOT NULL UNIQUE,
		"EncryptedPassword" TEXT NOT NULL);`
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

func createPostsTable(db *sql.DB) error {
	queryStr := `CREATE TABLE IF NOT EXISTS "posts" (
		"ID" INTEGER PRIMARY KEY AUTOINCREMENT,
		"AuthorID" INTEGER NOT NULL,
		"Author" TEXT NOT NULL,
		"Content" TEXT NOT NULL,
		"Categories" TEXT NOT NULL,
		"Timestamp" BIG INTEGER NOT NULL);`
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

func createCommentsTable(db *sql.DB) error {
	queryStr := `CREATE TABLE IF NOT EXISTS "comments" (
		"ID" INTEGER PRIMARY KEY AUTOINCREMENT,
		"PostID" INTEGER NOT NULL,
		"AuthorID" INTEGER NOT NULL,
		"Author" TEXT NOT NULL,
		"Content" TEXT NOT NULL,
		"Timestamp" BIG INTEGER NOT NULL);`
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
