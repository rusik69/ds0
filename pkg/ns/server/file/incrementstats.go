package file

import (
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/sirupsen/logrus"
)

// incrementStats increments the stats.
func incrementStats(fileSize uint64) error {
	logrus.Println("Increment stats")
	filesInfo, err := dbfile.GetFilesInfo()
	if err != nil {
		return err
	}
	logrus.Println("stats before increment: ", filesInfo.UncommittedSize, filesInfo.UncommittedFiles)
	filesInfo.UncommittedSize += fileSize
	filesInfo.UncommittedFiles++
	logrus.Println("stats after increment: ", filesInfo.UncommittedSize, filesInfo.UncommittedFiles)
	err = dbfile.SetFilesInfo(filesInfo)
	return err
}
