package event_db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("EVENT_DB_HOST")
	port     = os.Getenv("EVENT_DB_PORT")
	user     = os.Getenv("EVENT_DB_USER")
	password = os.Getenv("EVENT_DB_PASS")
	dbname   = os.Getenv("EVENT_DB_DBNAME")
)

var ErrFooConnectionError = errors.New("no se puede conectar a la eventdb PostgreSQL")
var eventDB *sql.DB

func Init() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	eventDB = db
	if err != nil {
		return ErrFooConnectionError
	}
	return nil
}

func GetEventDB() *sql.DB {
	return eventDB
}
