package io

import (
	"encoding/json"

	"github.com/foiovituh/hash-watcher/internal/util"
)

type Notification struct {
	Endpoint string `json:"endpoint"`
	Token    string `json:"token"`
}

func (n *Notification) Configured() bool {
	return n.Endpoint != "" && n.Token != ""
}

type WatcherSetting struct {
	DirectoryPath           string       `json:"directoryPath"`
	FileNames               []string     `json:"fileNames"`
	CheckFrequencyInSeconds int          `json:"checkFrequencyInSeconds"`
	Notification            Notification `json:"notification"`
}

func (ws *WatcherSetting) UnmarshalJson(filePath string) {
	fileBytes := ReadFile(filePath)
	err := json.Unmarshal(fileBytes, &ws)

	util.FatalIfTrue(err != nil, "Configuration file with invalid JSON!")
}
