package file

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// Download downloads a file.
func Download(src, dst, host, port string) error {
	url := "http://" + host + ":" + port + "/api/v0/file/download?file=" + src
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("download failed: " + url + " " + http.StatusText(resp.StatusCode))
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var fileInfo db.FileInfo
	if err := json.Unmarshal(bodyBytes, &fileInfo); err != nil {
		return err
	}
	success := false
	for _, node := range fileInfo.Nodes {
		url := "http://" + node.Host + ":" + node.Port + "/file/download?file=" + src
		resp, err := http.Get(url)
		if err != nil {
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			continue
		}
		file, err := CreateFile(src, dst)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := io.Copy(file, resp.Body); err != nil {
			return err
		}
		success = true
		break
	}
	if !success {
		return errors.New("download failed")
	}
	return nil
}
