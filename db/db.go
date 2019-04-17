package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

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
	    updated INTEGER
	)
	`

	_, err = db.Exec(sqlCreateTables)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateAsset stores new asset
func CreateAsset() int64 {
	statement := `INSERT INTO assets (success, created, updated) values (?, ?, ?)`

	conn.m.Lock()
	defer conn.m.Unlock()

	tx, err := conn.db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(statement)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	now := time.Now().UnixNano()

	_, err = stmt.Exec(1, now, now)
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

// Test tests db connection
func Test() {
	db := conn.db
	//defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()

	rows, err := db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err = db.Prepare("select name from foo where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow("3").Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	_, err = db.Exec("delete from foo")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("insert into foo(id, name) values(1, 'foo'), (2, 'bar'), (3, 'baz')")
	if err != nil {
		log.Fatal(err)
	}

	rows, err = db.Query("select id, name from foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
