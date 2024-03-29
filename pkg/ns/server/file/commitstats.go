package file

import (
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/sirupsen/logrus"
)

// commitStats commits the stats.
func commitStats(fileSize uint64) error {
	logrus.Println("Commit stats: ", "fileSize", fileSize)
	filesInfo, err := dbfile.GetFilesInfo()
	if err != nil {
		return err
	}
	filesInfo.UncommittedSize -= fileSize
	filesInfo.UncommittedFiles--
	filesInfo.TotalSize += fileSize
	filesInfo.TotalFiles++
	err = dbfile.SetStats(filesInfo)
	if err != nil {
		return err
	}
	return nil
}
