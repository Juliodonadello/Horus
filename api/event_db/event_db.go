package event_db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dashboard"
	password = "dashboard"
	dbname   = "events"
)

var connectionError = errors.New("No se puede conectar a la eventdb PostgreSQL")
var eventDB *sql.DB

func Init() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	eventDB, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return connectionError
	}
	defer eventDB.Close()
	return nil
}

func GetEventDB() *sql.DB {
	return eventDB
}
