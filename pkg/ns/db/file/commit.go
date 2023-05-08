package file

import "github.com/sirupsen/logrus"

// Commit commits the file to the database.
func Commit(fileName string) error {
	logrus.Println("Commit file: " + fileName)
	fileInfo, err := Get(fileName)
	if err != nil {
		return err
	}
	fileInfo.Committed = true
	err = Set(fileName, fileInfo)
	if err != nil {
		return err
	}
	return nil
}
