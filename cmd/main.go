package main

import (
	"os"

	"github.com/foiovituh/hash-watcher/internal/engine"
	"github.com/foiovituh/hash-watcher/internal/io"
	"github.com/foiovituh/hash-watcher/internal/util"
)

func main() {
	util.FatalIfTrue(len(os.Args) < 2, "Unspecified configuration file path!")

	watcherSetting := io.WatcherSetting{}
	watcherSetting.UnmarshalJson(os.Args[1])

	engine.Watch(&watcherSetting)
}
