package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var ACCESS_LOGS = os.Getenv("ACCESS_LOGS")
var PORT = os.Getenv("PORT")

func getDomainFromRequest(request *http.Request) string {
	domain := request.URL.Path[1:]
	return domain
}

func handleRequest(w http.ResponseWriter, request *http.Request) {
	if ACCESS_LOGS == "true" {
		fmt.Printf("/\n")
	}

	domain := getDomainFromRequest(request)
	log.Printf("Received new archive for domain '%s'\n", domain)

	tmpPath := saveFileFromRequestFormData(request)
	dstPath := filepath.Join(TMP_DIR, domain)
	unzipArchive(tmpPath, dstPath)
	defer deleteFile(tmpPath)

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK\n")
}

func main() {
	createInitialFiles()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRequest)

	fmt.Printf("Listening on 8080")
	http.ListenAndServe(":8080", mux)
}
