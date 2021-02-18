package main

import (
	"fmt"
	"time"

	"github.com/redochen/filecached/config"
)

func main() {
	startWatcher()
	fmt.Println("[main] watcher started!")

	for true {
		time.Sleep(1 * time.Second)
	}
}

func startWatcher() {
	if nil == config.Depository ||
		nil == config.Depository.Directories ||
		len(config.Depository.Directories) == 0 {
		return
	}

	for _, dir := range config.Depository.Directories {
		if len(dir.Path) == 0 || dir.Duration <= 0 {
			continue
		}

		go WatchDirectory(dir)
	}
}
