package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r, err := zip.OpenReader("old.apks")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	newFile, err := os.Create("stripped.apks")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	w := zip.NewWriter(newFile)
	defer w.Close()

	// Iterate through each file in the archive, appending them to a new ZIP archive unless
	// they're standalone APKs
	for _, f := range r.File {
		// We don't care about standalone APKs
		if strings.Contains(f.Name, "standalone") {
			continue
		}

		fileReader, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		fileWriter, err := w.Create(f.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(fileWriter, fileReader)
		if err != nil {
			log.Fatal(err)
		}
	}
}
