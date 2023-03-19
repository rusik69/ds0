package file

import (
	"github.com/sirupsen/logrus"
)

// Commit commits the file to the database.
func Commit(fileName string) error {
	fileInfo, err := Get(fileName)
	if err != nil {
		logrus.Error(err)
		return err
	}
	fileInfo.Committed = true
	err = Set(fileName, fileInfo)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
