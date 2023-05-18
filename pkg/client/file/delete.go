package file

import "net/http"

// Delete deletes the file.
func Delete(fileName, host, port string) error {
	url := "http://" + host + ":" + port + "/file/delete?file=" + fileName
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
