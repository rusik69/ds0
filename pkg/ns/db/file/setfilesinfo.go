package file

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// SetStats sets the files information.
func SetStats(filesInfo db.FilesInfo) error {
	logrus.Println("Set stats: ", "TotalFiles", filesInfo.TotalFiles,
		"TotalSize", filesInfo.TotalSize, "UncommittedFiles", filesInfo.UncommittedFiles,
		"UncommittedSize", filesInfo.UncommittedSize)
	filesInfoBytes, err := json.Marshal(filesInfo)
	if err != nil {
		logrus.Error(err)
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = db.DB.Put(ctx, "/stats", string(filesInfoBytes))
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
