package file

import (
	"net/http"
	"os"
)

// Upload uploads a file to the server.
func Upload(src, dst, host, port string) error {
	file, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	url := "http://" + host + ":" + port + "/upload?url=" + dst
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		panic("upload failed")
	}
	
	return nil
}
