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
	if filesInfo.UncommittedSize > fileSize {
		filesInfo.UncommittedSize -= fileSize
	}
	if filesInfo.UncommittedFiles > 0 {
		filesInfo.UncommittedFiles--
	}
	err = dbfile.SetFilesInfo(filesInfo)
	return err
}
