package main_test

import (
	"crypto/rand"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/rusik69/ds0/pkg/client/file"
)

var TestFileName string

// prepareFile prepares the test file.
func prepareFile() (string, error) {
	// Generate random data
	size := 1024 * 1024 * 10 // 10 MB
	randomData := make([]byte, size)
	_, err := rand.Read(randomData)
	if err != nil {
		return "", err
	}

	// Create a temporary file
	tempFile, err := ioutil.TempFile("", "ds0_test_*.tmp")
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	// Write random data to the temporary file
	_, err = tempFile.Write(randomData)
	if err != nil {
		return "", err
	}
	return tempFile.Name(), nil
}

// TestMain prepares the test data.
func TestMain(m *testing.M) {
	// Prepare the test file
	testFileName, err := prepareFile()
	if err != nil {
		log.Fatal(err)
	}
	TestFileName = testFileName
	code := m.Run()
	os.Exit(code)
}

// TestClient tests the client.
func TestClient(t *testing.T) {
	fileNameSplit := strings.Split(TestFileName, "/")
	fileName := fileNameSplit[len(fileNameSplit)-1]
	t.Run("upload", func(t *testing.T) {
		err := file.Upload(TestFileName, "/"+fileName, "ds0-ns", "6969")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("download", func(t *testing.T) {
		err := file.Download("/"+fileName, TestFileName, "ds0-ns", "6969")
		if err != nil {
			t.Error(err)
		}
	})
}
