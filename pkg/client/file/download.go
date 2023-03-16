package file

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// Download downloads a file.
func Download(src, dst, host, port string) {
	url := "http://" + host + ":" + port + "/download?url=" + src
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("download failed")
	}
	var nodes []db.HostInfo
	if err := json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
		panic(err)
	}
	success := false
	for _, node := range nodes {
		url := "http://" + node.Host + ":" + node.Port + "/download?url=" + src
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("download from %s has failed\n", node.Host)
			continue
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Printf("download from %s has failed\n", node.Host)
			continue
		}
		file, err := CreateFile(src, dst)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if _, err := io.Copy(file, resp.Body); err != nil {
			panic(err)
		}
		success = true
	}
	if !success {
		panic("download failed")
	}
}
