package filescan

import (
	"log"
	"os"

	"github.com/OPSWAT/mdcloud-go/api"
	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

// Scan or watches files or path
func Scan(api api.API, path string, watch bool) {
	fi, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		watchDir(api, path)
	case mode.IsRegular():
		api.ScanFile(path)
	}
}

func watchDir(api api.API, path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if (event.Op == fsnotify.Write) || (event.Op == fsnotify.Create) {
					api.ScanFile(event.Name)
				}
			case err := <-watcher.Errors:
				log.Fatal(err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
