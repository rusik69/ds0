package file

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// GetFile returns the node that stores the file.
func GetFile(fileName string) (db.FileInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := db.DB.Get(ctx, "/files/"+fileName)
	cancel()
	if err != nil {
		logrus.Fatal(err)
		return db.FileInfo{}, err
	}
	if len(resp.Kvs) == 0 {
		return db.FileInfo{}, errors.New("file not found")
	}
	var nodes []db.HostInfo
	for _, ev := range resp.Kvs {
		var s db.HostInfo
		json.Unmarshal(ev.Value, &s)
		nodes = append(nodes, s)
	}
	fileInfo := db.FileInfo{
		Nodes:     nodes,
		Committed: false,
	}
	return fileInfo, nil
}
