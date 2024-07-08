package io

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/foiovituh/hash-watcher/internal/util"
)

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)

	util.FatalIfError(err)

	return data
}

func GetFileData(directoryPath string, fileNames []string) map[string][]byte {
	fileDataPerFileName := make(map[string][]byte)

	if len(fileNames) == 0 {
		files, err := os.ReadDir(directoryPath)

		util.FatalIfError(err)

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			joinAndMap(directoryPath, file.Name(), fileDataPerFileName)
		}

		return fileDataPerFileName
	}

	for _, fileName := range fileNames {
		joinAndMap(directoryPath, fileName, fileDataPerFileName)
	}

	return fileDataPerFileName
}

func joinAndMap(directoryPath, fileName string, fileDataMap map[string][]byte) {
	filePath := joinDirectoryPathWithFileName(directoryPath, fileName)
	fileDataMap[fileName] = ReadFile(filePath)
}

func joinDirectoryPathWithFileName(directoryPath, filName string) string {
	separator := string(filepath.Separator)

	if strings.HasSuffix(directoryPath, separator) {
		separator = ""
	}

	return directoryPath + separator + filName
}
