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

// ScanOptions from scan cmd
type ScanOptions struct {
	Path         []string
	Headers      []string
	Watcher      bool
	Sanitization bool
	LookupFile   bool
	Poll         bool
}

var watcher *fsnotify.Watcher

// Scan or watches files or path
func Scan(api api.API, options ScanOptions) {
	if options.Path != nil && len(options.Path) > 0 {
		fi, err := os.Stat(options.Path[0])
		if err != nil {
			logrus.Fatalln(err)
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			watchScan(api, options)
		case mode.IsRegular():
			if !options.LookupFile {
				if res, err := api.ScanFile(options.Path[0], options.Headers, options.Poll); err == nil {
					logrus.Println(res)
				} else {
					logrus.Fatalln(err)
				}
			} else {
				lookupSHA1(api, options.Path[0], options)
			}
		}
	}
}

func watchScan(api api.API, options ScanOptions) {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(options.Path[0], addPaths); err != nil {
		logrus.Fatalln(err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if (event.Op == fsnotify.Write) || (event.Op == fsnotify.Create) {
					logrus.WithFields(logrus.Fields{"op": event.Op, "type": event.Name}).Infoln("Change detected")
					lookupSHA1(api, event.Name, options)
				}
			case err := <-watcher.Errors:
				logrus.Fatalln(err)
			}
		}
	}()
	<-done
}

func lookupSHA1(api api.API, filePath string, options ScanOptions) {
	resSha1, err := utils.GetFileSha1(filePath)
	if err != nil {
		api.ScanFile(filePath, options.Headers, options.Poll)
	} else {
		if res, err := api.FindOrScan(filePath, resSha1, options.Headers, options.Poll); err == nil {
			logrus.Println(res)
		} else {
			logrus.Fatalln(err)
		}
	}
}

func addPaths(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() && !strings.HasPrefix(fi.Name(), ".") {
		return watcher.Add(path)
	}
	return nil
}
