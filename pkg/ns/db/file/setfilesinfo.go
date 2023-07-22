package file

import (
	"context"
	"encoding/json"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// SetFilesInfo sets the files information.
func SetFilesInfo(filesInfo db.FilesInfo) error {
	logrus.Println("Set files info")
	filesInfoBytes, err := json.Marshal(filesInfo)
	if err != nil {
		logrus.Error(err)
		return err
	}
	logrus.Println("filesInfoBytes: ", string(filesInfoBytes))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = db.DB.Put(ctx, "/stats", string(filesInfoBytes))
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
