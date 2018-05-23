package filescan

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/OPSWAT/mdcloud-go/api"
	"github.com/OPSWAT/mdcloud-go/utils"
	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

// Scan or watches files or path
func Scan(api api.API, path, headers []string, watch, lookupFile bool) {
	if path != nil && len(path) > 0 {
		fi, err := os.Stat(path[0])
		if err != nil {
			logrus.Fatalln(err)
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			watchDirScan(api, path[0], headers)
		case mode.IsRegular():
			if !lookupFile {
				if res, err := api.ScanFile(path[0], headers); err == nil {
					logrus.Println(res)
				} else {
					logrus.Fatalln(err)
				}
			} else {
				lookupSha1(api, path[0], headers)
			}
		}
	}
}

func watchDirScan(api api.API, path string, headers []string) {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(path, watchDir); err != nil {
		logrus.Fatalln(err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if (event.Op == fsnotify.Write) || (event.Op == fsnotify.Create) {
					logrus.WithFields(logrus.Fields{"op": event.Op, "type": event.Name}).Infoln("Change detected")
					lookupSha1(api, event.Name, headers)
				}
			case err := <-watcher.Errors:
				logrus.Fatalln(err)
			}
		}
	}()
	<-done
}
func watchDir(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() && !strings.HasPrefix(fi.Name(), ".") {
		return watcher.Add(path)
	}
	return nil
}

func lookupSha1(api api.API, path string, headers []string) {
	resSha1, err := utils.GetFileSha1(path)
	if err != nil {
		api.ScanFile(path, headers)
	} else {
		if res, err := api.FindOrScan(path, resSha1, headers); err == nil {
			logrus.Println(res)
		} else {
			logrus.Fatalln(err)
		}
	}
}
