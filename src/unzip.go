package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// Unzips an archive
func unzipArchive(src string, dest string) {
	archive, err := zip.OpenReader(src)
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	os.MkdirAll(dest, os.ModePerm)

	for _, f := range archive.File {
		newPath := filepath.Join(dest, f.Name)
		log.Println(newPath)

		if f.FileInfo().IsDir() {
			os.MkdirAll(newPath, os.ModePerm)
			continue
		}

		dstFile, err := os.OpenFile(newPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}
		defer dstFile.Close()

		srcFile, err := f.Open()
		if err != nil {
			panic(err)
		}

		io.Copy(dstFile, srcFile)
		defer srcFile.Close()
	}

}
