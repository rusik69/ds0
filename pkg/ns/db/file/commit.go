package file

import "github.com/sirupsen/logrus"

// Commit commits the file to the database.
func Commit(fileName string) error {
	logrus.Println("Commit file: " + fileName)
	fileInfo, err := GetUncommitted(fileName)
	if err != nil {
		return err
	}
	fileInfo.Committed = true
	err = Set(fileName, fileInfo)
	if err != nil {
		return err
	}
	err = DeleteUncommitted(fileName)
	return err
}
