package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func includes(nameSlice []string, word string) bool {
	for _, name := range nameSlice {
		if strings.ToLower(name) == word {
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

	for _, item := range fileNames {

		splitName := strings.Split(item, "-")

		if includes(splitName, "frontend") {
			cmd := exec.Command(
				"osascript",
				"-e", fmt.Sprintf(`tell application "Terminal" to do script "cd %s/%s && yarn dev"`, rootDir, item),
			)

			if err := cmd.Start(); err != nil {
				log.Fatal(err.Error())
			} else {
				fmt.Println("Running Frontend service -> " + item)
			}
		} else {
			cmd2 := exec.Command(
				"osascript",
				"-e", fmt.Sprintf(`tell application "Terminal" to do script "cd %s/%s && mvn spring-boot:run"`, rootDir, item),
			)

			if err := cmd2.Start(); err != nil {
				log.Fatal(err.Error())
			} else {
				fmt.Println("Running Backend service -> " + item)
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
