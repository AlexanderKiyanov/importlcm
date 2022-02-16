package helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func UnzipFile(zipFile, destDir string) {
	archive, err := zip.OpenReader(zipFile)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = archive.Close()
		if err != nil {
			log.Fatalf("error: cant close archive %s", zipFile)
		}
	}()

	for _, f := range archive.File {
		filePath := filepath.Join(destDir, f.Name)
		fmt.Println("unzipping file ", filepath.Base(filePath))

		if !strings.HasPrefix(filePath, filepath.Clean(destDir)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			err = os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				panic(err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		err = dstFile.Close()
		if err != nil {
			panic(err)
		}
		err = fileInArchive.Close()
		if err != nil {
			panic(err)
		}
	}
}
