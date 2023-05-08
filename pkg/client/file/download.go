package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// Download downloads a file.
func Download(src, dst, host, port string) error {
	url := "http://" + host + ":" + port + "/file/download?file=" + src
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("download failed: " + url + " " + http.StatusText(resp.StatusCode))
	}
	var fileInfo db.FileInfo
	if err := json.NewDecoder(resp.Body).Decode(&fileInfo); err != nil {
		return err
	}
	success := false
	for _, node := range fileInfo.Nodes {
		url := "http://" + node.Host + ":" + node.Port + "/file/download?file=" + src
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("download failed: " + url)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("download failed: " + url + " " + http.StatusText(resp.StatusCode))
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
