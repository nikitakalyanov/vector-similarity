package vectorize

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

const createDataTableForTexts = "CREATE TABLE IF NOT EXISTS text_data(id serial, data text)"

const createExtensionPgVector = "CREATE EXTENSION IF NOT EXISTS vector"

func connectToDB() (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	return db, err
}

const insertTextData = "INSERT INTO text_data (data) VALUES ($1)"

type TextData struct {
	Id   int    `db:"id"`
	Data string `db:"data"`
}

const createTextEmbeddingTable = "CREATE TABLE IF NOT EXISTS text_embeddings (id int, embedding vector( 1536 ))"

const createTextVectorIdx = "CREATE INDEX ON text_embeddings USING ivfflat (embedding vector_l2_ops);"

const insertTextEmbedding = "INSERT INTO text_embeddings (id, embedding) VALUES ($1, $2::float4[])"

const getTextMatch = "SELECT id FROM text_embeddings ORDER BY embedding <-> $1::float4[]::vector LIMIT 1"
const selectTextData = "SELECT data FROM text_data WHERE id = $1"
