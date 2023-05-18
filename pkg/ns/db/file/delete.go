package file

import (
	"context"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// Delete removes the file from the database.
func Delete(fileName string) error {
	logrus.Println("Delete file: " + fileName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.DBFiles.Delete(ctx, "/files/"+fileName)
	cancel()
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
