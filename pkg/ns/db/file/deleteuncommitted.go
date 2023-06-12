package file

import (
	"context"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// DeleteUncommitted removes the file from the database.
func DeleteUncommitted(fileName string) error {
	logrus.Println("Delete uncommitted file: " + fileName)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	_, err := db.DB.Delete(ctx, "/files_uncommitted/"+fileName)
	cancel()
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
