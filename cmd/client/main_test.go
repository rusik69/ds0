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
var TestFilesNames []string

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

// prepareBenchmarkFiles prepares the benchmark files.
func prepareBenchmarkFiles() ([]string, error) {
	var res []string
	for i := 0; i < 1000; i++ {
		fileName, err := prepareFile()
		if err != nil {
			return nil, err
		}
		res = append(res, fileName)
	}
	return res, nil
}

// removeBenchmarkFiles removes the benchmark files.
func removeBenchmarkFiles() error {
	for _, fileName := range TestFilesNames {
		err := os.Remove(fileName)
		if err != nil {
			return err
		}
	}
	return nil
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
	testFilesNames, err := prepareBenchmarkFiles()
	if err != nil {
		log.Fatal(err)
	}
	TestFilesNames = testFilesNames
	waitForServer()
	code := m.Run()
	err = removeBenchmarkFiles()
	if err != nil {
		log.Fatal(err)
	}
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
		if len(nodes) != 5 {
			t.Error("expected 5 nodes, got ", len(nodes))
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
	t.Run("webmetrics", func(t *testing.T) {
		data, err := http.Get("http://ds0-web/metrics")
		if err != nil {
			t.Error(err)
		}
		if data.StatusCode != 200 {
			t.Error("expected 200")
		}
	})
	t.Run("nsmetrics", func(t *testing.T) {
		data, err := http.Get("http://ds0-ns:6969/metrics")
		if err != nil {
			t.Error(err)
		}
		if data.StatusCode != 200 {
			t.Error("expected 200")
		}
	})
	t.Run("nodemetrics", func(t *testing.T) {
		nodes, err := node.List("ds0-ns", "6969")
		if err != nil {
			t.Error(err)
		}
		for _, node := range nodes {
			data, err := http.Get("http://" + node.Host + ":" + node.Port + "/metrics")
			if err != nil {
				t.Error(err)
			}
			if data.StatusCode != 200 {
				t.Error("expected 200")
			}
		}
	})
}
