package engine

import (
	"fmt"
	"log"
	"os"

	"github.com/foiovituh/hash-watcher/internal/io"
	"github.com/foiovituh/hash-watcher/internal/util"
	"github.com/slack-go/slack"
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
		message := fmt.Sprint("=> ", file, " was modified!",
			"\n  - Before: ", hh.PreviousHash,
			"\n  - Now: ", hh.CurrentHash)

		Logger.Print(message)

		if ws.Notification.Configured() {
			SendToSlackChannel(
				ws.Notification.Endpoint,
				ws.Notification.Token,
				message,
			)
		}

		hh.Notified = true
	}
}

func SendToSlackChannel(channelID, token, message string) {
	client := slack.New(token)

	channelID, _, err := client.PostMessage(
		channelID,
		slack.MsgOptionText(message, false),
	)

	util.FatalIfError(err)
	log.Println("=> Message sent to the Slack channel:", channelID)
}
