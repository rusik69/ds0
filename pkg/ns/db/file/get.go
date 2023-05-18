package file

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// Get returns the node that stores the file.
func Get(fileName string) (db.FileInfo, error) {
	logrus.Println("Get file: " + fileName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := db.DBFiles.Get(ctx, "/files/"+fileName)
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
