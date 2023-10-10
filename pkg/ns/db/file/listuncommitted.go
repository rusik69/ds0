package file

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// ListUncommitted returns the uncommitted files.
func ListUncommitted() (map[string]db.FileInfo, error) {
	//logrus.Println("List uncommitted files")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := db.DB.Get(ctx, "/files_uncommitted")
	cancel()
	if err != nil {
		return nil, err
	}
	res := map[string]db.FileInfo{}
	if len(resp.Kvs) == 0 {
		return res, errors.New("no uncommitted files found")
	}
	for _, kv := range resp.Kvs {
		var fileInfo db.FileInfo
		err := json.Unmarshal(kv.Value, &fileInfo)
		if err != nil {
			continue
		}
		res[string(kv.Key)] = fileInfo
	}
	return res, nil
}
