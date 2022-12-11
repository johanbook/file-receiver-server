package main

import (
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var ROOT_FILE_PATH = os.Getenv("ROOT_FILE_PATH")

const TMP_DIR = "/tmp/file-receiver-server/"

func createInitialFiles() {
	_, err := os.Stat(TMP_DIR)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(TMP_DIR, 0755)

		if errDir != nil {
			log.Fatal(err)
		}
	}
}

// Deletes a file from host sytem
func deleteFile(src string) {
	os.Remove(src)
}

// Creates a path for a temporary file
func getTempFilePath() string {
	id := uuid.New().String()
	return filepath.Join(TMP_DIR, id)
}

func saveFileFromRequestFormData(request *http.Request) string {
	request.ParseMultipartForm(10 << 20)
	file, _, err := request.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	tempFilePath := getTempFilePath()
	out, err := os.Create(tempFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	io.Copy(out, file)

	return tempFilePath

}
