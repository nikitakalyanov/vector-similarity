package vectorize

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/nikitakalyanov/vector-similarity/pkg/openaiclient"
)

const imageKind = "image"
const textKind = "text"
const vectorEngineOpenAI = "openai"

func DoInit(kind string) error {
	switch kind {
	case textKind:
		db, err := connectToDB()
		if err != nil {
			return err
		}
		defer db.Close(context.TODO())
		_, err = db.Exec(context.TODO(), createExtensionPgVector)
		if err != nil {
			return err
		}
		_, err = db.Exec(context.TODO(), createDataTableForTexts)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("unknown kind")
	}
}

func DoSeed(kind string, inputDir string) error {
	switch kind {
	case textKind:
		return processTextDir(inputDir)
	default:
		return errors.New("unknown kind")
	}
}

func findTextMatch(inputFile string) error {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}
	// vectorize here
	openAI := openaiclient.NewClient()
	res, err := openAI.GetEmbeddings(context.TODO(), string(data))
	if err != nil {
		return err
	}
	db, err := connectToDB()
	if err != nil {
		return err
	}
	rows, err := db.Query(context.TODO(), getTextMatch, res)
	if err != nil {
		return err
	}
	db2, err := connectToDB()
	if err != nil {
		return err
	}
	for rows.Next() {
		vals, err := rows.Values()
		if err != nil {
			return err
		}
		row := db2.QueryRow(context.TODO(), selectTextData, vals[0])
		var data string
		err = row.Scan(&data)
		if err != nil {
			return err
		}
		fmt.Println("found closest match:")
		fmt.Println(data)
	}

	return nil
}

func DoFind(kind string, inputFile string) error {
	switch kind {
	case textKind:
		return findTextMatch(inputFile)
	default:
		return errors.New("unknown kind")
	}
}

func DoVectorize(kind string, vectorEngine string) error {
	switch kind {
	case textKind:
		return vectorizeText(vectorEngine)
	default:
		return errors.New("unknown kind")
	}
}

func vectorizeTextWithOpenAI() error {
	db, err := connectToDB()
	if err != nil {
		return err
	}
	defer db.Close(context.TODO())

	_, err = db.Exec(context.TODO(), createTextEmbeddingTable)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = db.Exec(context.TODO(), createTextVectorIdx)
	if err != nil {
		return err
	}

	rows, err := db.Query(context.TODO(), "SELECT * FROM text_data")
	if err != nil {
		return err
	}
	db2, err := connectToDB()
	if err != nil {
		return err
	}
	defer db2.Close(context.TODO())
	for rows.Next() {
		vals, err := rows.Values()
		if err != nil {
			return err
		}
		// vectorize here
		openAI := openaiclient.NewClient()
		res, err := openAI.GetEmbeddings(context.TODO(), vals[1].(string))
		if err != nil {
			return err
		}
		_, err = db2.Exec(context.TODO(), insertTextEmbedding, vals[0], res)
		if err != nil {
			return err
		}
	}

	return nil
}

func vectorizeText(vectorEngine string) error {
	switch vectorEngine {
	case vectorEngineOpenAI:
		return vectorizeTextWithOpenAI()
	default:
		return errors.New("unknown vectorizetion engine")
	}
}
