package main_test

import (
	"crypto/rand"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/rusik69/ds0/pkg/client/file"
	"github.com/rusik69/ds0/pkg/client/node"
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
	tempFile, err := os.CreateTemp("", "ds0_test_*.tmp")
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

// waitForServer waits for the server to start.
func waitForServer() {
	for {
		_, err := http.Get("http://ds0-ns:6969/ping")
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// TestMain prepares the test data.
func TestMain(m *testing.M) {
	// Prepare the test file
	testFileName, err := prepareFile()
	if err != nil {
		log.Fatal(err)
	}
	TestFileName = testFileName
	waitForServer()
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
	t.Run("delete", func(t *testing.T) {
		err := file.Delete("/"+fileName, "ds0-ns", "6969")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("listnodes", func(t *testing.T) {
		nodes, err := node.List("ds0-ns", "6969")
		if err != nil {
			t.Error(err)
		}
		if len(nodes) != 3 {
			t.Error("expected 3 nodes")
		}
	})
	t.Run("webroot", func(t *testing.T) {
		data, err := http.Get("http://ds0-web/")
		if err != nil {
			t.Error(err)
		}
		if data.StatusCode != 200 {
			t.Error("expected 200")
		}
	})
}
