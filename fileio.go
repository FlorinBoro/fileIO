package fileio

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"gopkg.in/yaml.v3"
)

type globalSettings struct {
	defaultFilePermissions fs.FileMode `yaml:"file permissions"`
	logsFolder             string      `yaml:"logs folder"`
	filesFolder            string      `yaml:"files folder"`
}

var Configs globalSettings

func simplePathCreate(path string) {
	os.MkdirAll(path, Configs.defaultFilePermissions)
}

func simpleFileCreate(path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Unable to create file due to above error...")
		return
	}
	defer file.Close()
	fmt.Println("File created successfully: ", path)
}

func createFolderPaths(checkIfExists bool, path string) {
	switch checkIfExists {
	case true:
		if _, err := os.Stat(path); os.IsNotExist(err) {
			os.MkdirAll(path, Configs.defaultFilePermissions)
		} else {
			fmt.Print("For path: ")
			fmt.Print(path)
			fmt.Println("Encountered the following error upon creating folder/path")
			fmt.Println(err)
			return
		}
	case false:
		os.MkdirAll(path, Configs.defaultFilePermissions)
		fmt.Println("Command to create path succesfully sent")
	}

}

func createFileAndPath(checkIfExists bool, path, fileName string) {
	createFolderPaths(checkIfExists, path)
	newPath := path + fileName
	if !isFileAlreadyThere(newPath) {
		file, err := os.Create(newPath)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Unable to create file due to above error...")
			return
		}
		defer file.Close()
		fmt.Println("File created successfully: ", newPath)
	} else {
		fmt.Println("File already there... <createFile> will take no further action")
	}
}

func isFileAlreadyThere(filePath string) bool {
	_, error := os.Stat(filePath)
	return !errors.Is(error, os.ErrNotExist)
}

func main() {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Configs)
	if err != nil {
		panic(err)
	}

	createFolderPaths(true, Configs.filesFolder)
}
