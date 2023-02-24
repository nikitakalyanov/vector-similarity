package vectorize

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const createDataTableForTexts = "CREATE TABLE IF NOT EXISTS text_data(id serial, data text)"

const createExtensionPgVector = "CREATE EXTENSION vector"

func connectToDB() (*sqlx.DB, error) {
	url, _ := os.LookupEnv("DATABASE_URL")
	db, err := sqlx.Connect("postgres", url)
	return db, err
}
