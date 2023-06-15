package file

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// GetUncommitted returns the node that stores the file.
func GetUncommitted(fileName string) (db.FileInfo, error) {
	logrus.Println("Get uncommitted file: " + fileName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := db.DB.Get(ctx, "/files_uncommitted/"+fileName)
	cancel()
	if err != nil {
		logrus.Error(err)
		return db.FileInfo{}, err
	}
	if len(resp.Kvs) == 0 {
		return db.FileInfo{}, os.ErrNotExist
	}
	var fileInfo db.FileInfo
	json.Unmarshal(resp.Kvs[0].Value, &fileInfo)
	logrus.Println("Get file: " + string(resp.Kvs[0].Value))
	return fileInfo, nil
}
