package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func MakeFullPath(path string, currDir string) (string, error) {
	runes := []rune(path)
	firstSymbol := string(runes[0:1])

	if IsDirectory(path) || IsFile(path) {
		return path, nil
	} else if firstSymbol != "\\" {
		path = fmt.Sprint(currDir, "\\", path)
	} else {
		path = fmt.Sprint(currDir, path)
	}

	if IsDirectory(path) || IsFile(path) {
		return path, nil
	}
	return path, errors.New("error: unrecognized arguments")
}

func IsDirectory(path string) bool {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		return true
	} else {
		return false
	}
}

func IsFile(path string) bool {

	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	if fileInfo.Mode().IsRegular() {
		return true
	} else {
		return false
	}
}

func FindFilesByPath(path, pattern string) ([]string, error) {

	var matches []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func ReplaceString(fileName, strFrom, strTo string) error {
	input, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, strFrom) {
			lines[i] = strTo
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(fileName, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
