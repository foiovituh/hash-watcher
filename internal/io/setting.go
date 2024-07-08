package io

import (
	"encoding/json"

	"github.com/foiovituh/hash-watcher/internal/util"
)

type WatcherSetting struct {
	DirectoryPath           string   `json:"directoryPath"`
	FileNames               []string `json:"fileNames"`
	CheckFrequencyInSeconds int      `json:"checkFrequencyInSeconds"`
	/*
		TODO:
		ExecutionTimeLimitInMinutes int    `json:"executionTimeLimitInMinutes"`
		NotificationType            string `json:"notificationType"`
		NotificationCredentials     ?      `json:"notificationCredentials`
	*/
}

func (ws *WatcherSetting) UnmarshalJson(filePath string) {
	fileBytes := ReadFile(filePath)
	err := json.Unmarshal(fileBytes, &ws)

	util.FatalIfTrue(err != nil, "Configuration file with invalid JSON!")
}
