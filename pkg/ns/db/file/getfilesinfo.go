package file

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// GetFilesInfo gets the clusters stats.
func GetFilesInfo() (db.FilesInfo, error) {
	logrus.Println("Get files info")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := db.DB.Get(ctx, "/stats")
	cancel()
	if err != nil {
		logrus.Error(err)
		return db.FilesInfo{}, err
	}
	if len(resp.Kvs) == 0 {
		// create empty stats
		filesInfoBytes, err := json.Marshal(db.FilesInfo{})
		if err != nil {
			logrus.Error(err)
			return db.FilesInfo{}, err
		}
		logrus.Println("filesInfoBytes: ", string(filesInfoBytes))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_, err = db.DB.Put(ctx, "/stats", string(filesInfoBytes))
		if err != nil {
			logrus.Error(err)
			return db.FilesInfo{}, err
		}
		logrus.Println("Empty stats created")
		return db.FilesInfo{}, nil
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		resp, err := db.DB.Get(ctx, "/stats")
		if err != nil {
			logrus.Error(err)
			return db.FilesInfo{}, err
		}
		logrus.Println("resp.Kvs[0].Value: ", string(resp.Kvs[0].Value))
		var filesInfo db.FilesInfo
		json.Unmarshal(resp.Kvs[0].Value, &filesInfo)
		logrus.Println("Stats: ", filesInfo.UncommittedSize, filesInfo.UncommittedFiles)
		return filesInfo, nil
	}
}
