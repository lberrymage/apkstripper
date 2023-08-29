package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: apkstripper <oldfile> <newfile>")
		os.Exit(1)
	}
	oldFileName := os.Args[1]
	newFileName := os.Args[2]

	r, err := zip.OpenReader(oldFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	newFile, err := os.Create(newFileName)
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
