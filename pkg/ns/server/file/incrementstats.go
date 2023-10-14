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
	filesInfo.UncommittedSize += fileSize
	filesInfo.UncommittedFiles++
	err = dbfile.SetStats(filesInfo)
	return err
}
