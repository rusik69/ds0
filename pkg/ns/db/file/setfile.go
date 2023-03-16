package file

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// SetFile sets file state in the dabase.
func Set(fileName string, fileInfo db.FileInfo) error {
	fileInfoBytes, err := json.Marshal(fileInfo)
	if err != nil {
		logrus.Error(err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = db.DB.Put(ctx, "/files/"+fileName, string(fileInfoBytes))
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	return nil
}
