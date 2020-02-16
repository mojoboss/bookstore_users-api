package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	PG_HOST     = "USER_API_PG_DB_HOST"
	PG_PORT     = "USER_API_PG_DB_PORT"
	PG_USER     = "USER_API_PG_DB_USER"
	PG_PASSWORD = "USER_API_PG_DB_PASSWORD"
	PG_DBNAME   = "USER_API_PG_DB_NAME"
	PG_SCHEMA   = "USER_API_PG_DB_SCHEMA"
)

var (
	host     = os.Getenv(PG_HOST)
	port     = os.Getenv(PG_PORT)
	user     = os.Getenv(PG_USER)
	password = os.Getenv(PG_PASSWORD)
	dbname   = os.Getenv(PG_DBNAME)
	schema   = os.Getenv(PG_SCHEMA)
)

var (
	Client *sql.DB
)

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s search_path=%s sslmode=disable", host, port, user, password, dbname, schema)
	Client, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
