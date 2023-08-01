package file

import (
	dbfile "github.com/rusik69/ds0/pkg/ns/db/file"
	"github.com/sirupsen/logrus"
)

// commitStats commits the stats.
func commitStats(fileSize uint64) error {
	logrus.Println("Commit stats")
	logrus.Println("fileSize: ", fileSize)
	filesInfo, err := dbfile.GetFilesInfo()
	if err != nil {
		return err
	}
	filesInfo.UncommittedSize -= fileSize
	filesInfo.UncommittedFiles--
	logrus.Println("filesInfo.UncommittedFiles: ", filesInfo.UncommittedFiles)
	filesInfo.TotalSize += fileSize
	logrus.Println("filesInfo.TotalSize: ", filesInfo.TotalSize)
	filesInfo.TotalFiles++
	logrus.Println("filesInfo.TotalFiles: ", filesInfo.TotalFiles)
	err = dbfile.SetFilesInfo(filesInfo)
	if err != nil {
		return err
	}
	return nil
}
