package file

import (
	"fmt"
	"os"
	"strings"
)

// CreateFile creates a file.
func CreateFile(src, dst string) (*os.File, error) {
	fullPath := ""
	if dst[len(dst)-1] == '/' {
		fullPath = dst + strings.Split(src, "/")[len(strings.Split(src, "/"))-1]
	} else {
		fullPath = dst
	}
	fmt.Println("Create file: " + fullPath)
	_, err := os.Stat(fullPath)
	if err != nil {
		if err == os.ErrNotExist {
			file, err := os.Create(fullPath)
			if err != nil {
				fmt.Println("Create file error: " + err.Error())
				return nil, err
			}
			return file, nil
		}
	} else {
		return os.Open(fullPath)
	}
	return nil, nil
}
