package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func includes(sliceOfString []string, wordToBeMatched string) bool {
	for _, name := range sliceOfString {
		if strings.ToLower(name) == wordToBeMatched {
			return true
		}
	}
	return false
}

func wakeUpServices(rootDir string) {

	files, err := os.ReadDir(rootDir)

	if err != nil {
		log.Fatal(err.Error())
	}

	var fileNames []string

	for _, file := range files {
		if file.IsDir() {
			fileNameSlice := strings.Split(file.Name(), "-")
			if includes(fileNameSlice, "ahm") {
				fileNames = append(fileNames, file.Name())
			}
		}
	}

	for _, dirName := range fileNames {

		sliceOfName := strings.Split(dirName, "-")

		if includes(sliceOfName, "frontend") {
			cmd := exec.Command(
				"osascript",
				"-e", fmt.Sprintf(`tell application "Terminal" to do script "cd %s/%s && yarn dev"`, rootDir, dirName),
			)

			if err := cmd.Start(); err != nil {
				log.Fatal(err.Error())
			} else {
				fmt.Println("Running Frontend service -> " + dirName)
			}
		} else {
			cmd2 := exec.Command(
				"osascript",
				"-e", fmt.Sprintf(`tell application "Terminal" to do script "cd %s/%s && mvn spring-boot:run"`, rootDir, dirName),
			)

			if err := cmd2.Start(); err != nil {
				log.Fatal(err.Error())
			} else {
				fmt.Println("Running Backend service -> " + dirName)
			}
		}
	}
}

func main() {

	fmt.Println("Checking for AHM_ROOT_DIR value ...")

	ahmRootDir, isFrontendDirPresent := os.LookupEnv("AHM_ROOT_DIR")

	if !isFrontendDirPresent {
		fmt.Println("No value found for AHM_ROOT_DIR")
		fmt.Println("exitting...")
		os.Exit(1)
	}
	fmt.Println("AHM_ROOT_DIR found to be ", ahmRootDir)

	wakeUpServices(ahmRootDir)
}
