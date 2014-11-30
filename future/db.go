package future

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/russross/meddler"
)

/* */
func LoadFutureDB() *sql.DB {
	db, err := sql.Open("sqlite3", "futuredb.db")
	if err != nil {
		panic("Error loading SQL DB: " + err.Error())
	}
	initDb(db)
	return db
}

func Get(id string) (*Future, error) {
	db := LoadFutureDB()
	future := new(Future)
	err := meddler.QueryRow(db, future, "SELECT * FROM future WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return future, err
}

/* */
func initDb(db *sql.DB) {
	if _, err := db.Exec(FUTURE_SCHEMA); err != nil {
		panic("error creating person table: " + err.Error())
	}
}
