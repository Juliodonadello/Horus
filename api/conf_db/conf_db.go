package conf_db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var (
	host     = os.Getenv("CONF_DB_HOST")
	port     = os.Getenv("CONF_DB_PORT")
	user     = os.Getenv("CONF_DB_USER")
	password = os.Getenv("CONF_DB_PASS")
	dbname   = os.Getenv("CONF_DB_DBNAME")
)

var connectionError = errors.New("No se puede conectar a la confdb PostgreSQL")
var confDB *sql.DB

func Init() error {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	confDB = db
	if err != nil {
		return connectionError
	}
	return nil
}

func GetConfDB() *sql.DB {
	return confDB
}
