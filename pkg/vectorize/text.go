package vectorize

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
)

func processTextDir(dir string) error {
	db, err := connectToDB()
	if err != nil {
		return err
	}
	defer db.Close(context.TODO())
	var walker fs.WalkDirFunc = func(
		path string, d fs.DirEntry, err error) error {
		if path == dir {
			return nil
		}
		if d.IsDir() && path != dir {
			return fs.SkipDir
		}
		if err != nil {
			return err
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		_, err = db.Exec(context.TODO(), insertTextData, string(data))
		return err
	}
	return filepath.WalkDir(dir, walker)
}
