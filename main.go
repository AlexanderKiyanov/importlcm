package main

import (
	"fmt"
	"importlcm/helpers"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path, err := helpers.GetOptions(os.Args[1:], currentDir)
	if err != nil {
		log.Fatal(err)
	}

	fullNameZipFiles, err := helpers.FindFilesByPath(path, "*.zip")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Elevate Privileges to Administrator level
	if !helpers.AmIAdmin() {
		helpers.ElevateAsAdmin()
	}

	utility := "C:\\Middleware\\user_projects\\epmsystem1\\bin\\Utility.bat"

	logFileName := fmt.Sprintf("epmpi_%s.log", time.Now().Format("20060102_150405"))
	logFile := filepath.Join(currentDir, logFileName)
	//logFile := fmt.Sprintf("%s\\epmpi_%s.log", currentDir, time.Now().Format("20060102_150405"))
	fmt.Printf("\n\nLog file is located: %s\n\n", logFile)

	for i := range fullNameZipFiles {
		fullNameZipFileNoExt := strings.TrimSuffix(fullNameZipFiles[i], filepath.Ext(fullNameZipFiles[i]))
		zipFileNoExt := strings.TrimSuffix(filepath.Base(fullNameZipFiles[i]), filepath.Ext(fullNameZipFiles[i]))

		tmpFileDir := "C:\\Users\\a_kiyanov\\Desktop\\temp"
		err := os.MkdirAll(tmpFileDir, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		tmpFileDir = filepath.Join(tmpFileDir, zipFileNoExt)

		//err = os.RemoveAll(tmpFileDir)
		//if err != nil {
		//	log.Fatal(err)
		//}

		helpers.UnzipFile(fullNameZipFiles[i], tmpFileDir)

		importXml := filepath.Join(tmpFileDir, "Import.xml")
		passString := "    <User name=\"\" password=\"\"/>"
		passStringNew := "    <User name=\"admin\" password=\"{LCM}i6ClGQGnlyFy/5LWpIJGB7gWuYPop+DfHKQMB/3kV1i1E+ZGgq7lfiEL1FepI+YB\"/>"

		if err != helpers.ReplaceString(importXml, passString, passStringNew) {
			err = os.RemoveAll(fullNameZipFileNoExt)
			if err != nil {
				log.Fatal(err)
			}

			log.Fatalf("error: can not open the Import.xml file :\n%s\n%s", importXml, err)
		}

		helpers.StartImport(utility, importXml, logFile)

		err = os.RemoveAll(fullNameZipFileNoExt)
		if err != nil {
			log.Fatal(err)
		}

	}

}
