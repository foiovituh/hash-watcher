package engine

import (
	"log"
	"os"

	"github.com/foiovituh/hash-watcher/internal/io"
)

var Logger *log.Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

type HashHistory struct {
	PreviousHash string
	CurrentHash  string
	Notified     bool
}

func (hh *HashHistory) Update(newHash string) {
	hh.PreviousHash = hh.CurrentHash
	hh.CurrentHash = newHash
	hh.Notified = false
}

func (hh *HashHistory) NotifyIfNecessary(ws *io.WatcherSetting, file string) {
	if hh.PreviousHash != hh.CurrentHash && !hh.Notified {

		Logger.Println("=>", file, "was modified!",
			"\n  - Before:", hh.PreviousHash,
			"\n  - Now:", hh.CurrentHash)

		hh.Notified = true
	}
}
