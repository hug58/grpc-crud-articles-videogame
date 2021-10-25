package data

import (
	"database/sql"
	"log"
	"sync"
)

var (
	data *Data
	once sync.Once
)

type Data struct {
	DB *sql.DB
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	err = MakeMigration(db)
	if err != nil {
		log.Panic(err)
	}

	data = &Data{
		DB: db,
	}
}

func New() *Data {
	once.Do(initDB)
	return data
}

// Close closes the resources used by data.
func Close() error {
	if data == nil {
		return nil
	}

	return data.DB.Close()
}
