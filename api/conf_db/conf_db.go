package conf_db

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
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
var confDB *sqlx.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlInfo)
	confDB = db
	if err != nil {
		panic(connectionError)
	}
}

func GetConfDB() *sqlx.DB {
	return confDB
}
