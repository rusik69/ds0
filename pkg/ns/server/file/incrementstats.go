package file

import dbfile "github.com/rusik69/ds0/pkg/ns/db/file"

// incrementStats increments the stats.
func incrementStats(fileSize uint64) error {
	filesInfo, err := dbfile.GetFilesInfo()
	if err != nil {
		return err
	}
	filesInfo.UncommittedSize += fileSize
	filesInfo.UncommittedFiles++
	err = dbfile.SetFilesInfo(filesInfo)
	return err
}
