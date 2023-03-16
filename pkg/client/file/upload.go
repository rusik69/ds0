package file

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rusik69/ds0/pkg/ns/db"
)

// Upload uploads a file to the server.
func Upload(src, dst, host, port string) error {
	url := "http://" + host + ":" + port + "/upload?url=" + dst
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("upload failed")
	}
	var nodes []db.HostInfo
	if err := json.NewDecoder(resp.Body).Decode(&nodes); err != nil {
		panic(err)
	}
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for _, node := range nodes {
		url := "http://" + node.Host + ":" + node.Port + "/upload?url=" + dst
		resp, err := http.Post(url, "application/octet-stream", file)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			panic("upload failed")
		}
	}
	return nil
}
