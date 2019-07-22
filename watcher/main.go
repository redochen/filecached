package main

import (
	"fmt"
	cfg "github.com/redochen/filecached/config"
	"time"
)

func main() {
	startWatcher()
	fmt.Println("[main] watcher started!")

	for true {
		time.Sleep(1 * time.Second)
	}
}

func startWatcher() {
	if nil == cfg.Depository ||
		nil == cfg.Depository.Directories ||
		len(cfg.Depository.Directories) == 0 {
		return
	}

	for _, dir := range cfg.Depository.Directories {
		if len(dir.Path) == 0 || dir.Duration <= 0 {
			continue
		}

		go WatchDirectory(dir)
	}
}
