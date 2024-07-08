package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/foiovituh/hash-watcher/internal/io"
)

func Watch(ws *io.WatcherSetting) {
	hashHistoryMap := make(map[string]*HashHistory)

	tickerDuration := time.Duration(ws.CheckFrequencyInSeconds)
	ticker := time.NewTicker(time.Second * tickerDuration)
	defer ticker.Stop()

	Logger.Print("=> Watching...\n\n")

	for range ticker.C {
		fileDataMap := io.GetFileData(ws.DirectoryPath, ws.FileNames)

		for fileName, fileData := range fileDataMap {
			hashed := sha256.Sum256(fileData)
			currentHash := hex.EncodeToString(hashed[:])

			if _, mapped := hashHistoryMap[fileName]; !mapped {
				hashHistoryMap[fileName] = &HashHistory{
					PreviousHash: currentHash,
					CurrentHash:  currentHash,
					Notified:     false,
				}

				continue
			}

			hashHistory := hashHistoryMap[fileName]
			hashHistory.Update(currentHash)
			hashHistory.NotifyIfNecessary(ws, fileName)
		}
	}
}
