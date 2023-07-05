package file

import (
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/sirupsen/logrus"
)

// decrementStats increments the stats.
func decrementStats(fileSize uint64) error {
	logrus.Println("Decrement stats")
	filesInfo, err := dbfile.GetFilesInfo()
	if err != nil {
		return err
	}
	logrus.Println("stats before decrement: ", filesInfo.UncommittedSize, filesInfo.UncommittedFiles)
	filesInfo.UncommittedSize -= fileSize
	filesInfo.UncommittedFiles--
	logrus.Println("stats after decrement: ", filesInfo.UncommittedSize, filesInfo.UncommittedFiles)
	err = dbfile.SetFilesInfo(filesInfo)
	return err
}
