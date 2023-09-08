package file

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/rusik69/ds0/pkg/ns/db"
	"github.com/sirupsen/logrus"
)

// List returns the list of files by prefix.
func List(prefix string) (map[string]db.FileInfo, error) {
	logrus.Println("List files in prefix: " + prefix)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := db.DB.Get(ctx, prefix)
	cancel()
	if err != nil {
		return nil, err
	}
	res := map[string]db.FileInfo{}
	if len(resp.Kvs) == 0 {
		return res, errors.New("no files found")
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
