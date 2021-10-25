package data

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	//uri := "postgres://hug58:grosss213@127.0.0.1:5432/microvideogame?sslmode=disable"
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)
	return sql.Open("postgres", pgConnString)
}

func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
