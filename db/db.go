package db

import (
	"database/sql"
	"log"
	"sync"

	// dumb go import
	_ "github.com/mattn/go-sqlite3"
)

// Connection serialized connection
type Connection struct {
	m  sync.Mutex
	db *sql.DB
}

var conn *Connection

func init() {
	//os.Remove("./sync.db")
	db, err := sql.Open("sqlite3", "./sync.db")

	if err != nil {
		log.Fatal(err)
	}

	conn = new(Connection)
	conn.db = db

	sqlCreateTables := `
	CREATE TABLE IF NOT EXISTS assets (
	    id INTEGER NOT NULL PRIMARY KEY,
	    success INTEGER,
	    created INTEGER,
	    updated INTEGER,
	    filename TEXT
	)
	`

	_, err = db.Exec(sqlCreateTables)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateAsset stores new asset
func CreateAsset(timestamp int64) int64 {
	conn.m.Lock()
	defer conn.m.Unlock()

	statement := `INSERT INTO assets (success, created, updated) values (?, ?, ?)`

	tx, err := conn.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(statement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(0, timestamp, timestamp)
	if err != nil {
		log.Fatal(err)
	}

	stmt2, err := tx.Prepare("SELECT last_insert_rowid()")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt2.Close()

	var id int64

	err = stmt2.QueryRow().Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	tx.Commit()

	return id
}

func SaveAsset(timestamp int64, id int64, filename string) {
	conn.m.Lock()
	defer conn.m.Unlock()

	statement := `UPDATE assets SET success = 1, updated = ?, filename = ? WHERE id = ?`

	stmt, err := conn.db.Prepare(statement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(timestamp, filename, id)
	if err != nil {
		log.Fatal(err)
	}
}
