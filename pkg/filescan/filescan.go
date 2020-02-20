package filescan

import (
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/OPSWAT/mdcloud-go/pkg/api"
	"github.com/OPSWAT/mdcloud-go/pkg/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
)

// ScanOptions from scan cmd
type ScanOptions struct {
	Path         []string
	Headers      []string
	Watcher      bool
	Sanitization bool
	Unarchive    bool
	LookupFile   bool
	Poll         bool
}

var watcher *fsnotify.Watcher

// Scan or watches files or path
func Scan(api *api.API, options ScanOptions) {
	if options.Path != nil && len(options.Path) > 0 {
		fname := options.Path[0]
		if !path.IsAbs(fname) {
			wd, err := os.Getwd()
			if err != nil {
				logrus.Panicln(err)
			}
			fname = path.Clean(path.Join(wd, fname))
			options.Path[0] = fname
		}
		fi, err := os.Stat(fname)
		if err != nil {
			logrus.Fatalln(err)
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			watchScan(api, options)
		case mode.IsRegular():
			if !options.LookupFile {
				if res, err := api.ScanFile(fname, options.Headers, options.Poll); err == nil {
					logrus.WithField("result", res).Infoln("Scan result")
				} else {
					logrus.Fatalln(err)
				}
			} else {
				lookupSHA1(api, fname, options)
			}
		}
	}
}

func watchScan(api *api.API, options ScanOptions) {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(options.Path[0], addPaths); err != nil {
		logrus.Fatalln(err)
	}

	var rate time.Duration
	if api.Type > 0 {
		rate = time.Minute / 10
	} else {
		rate = time.Minute / 100
	}
	throttle := time.Tick(rate)

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if !strings.Contains(event.Name, "/.") && event.Op == fsnotify.Write {
					logrus.WithFields(logrus.Fields{"op": event.Op, "type": event.Name}).Infoln("Change detected")
					<-throttle
					if len(api.Limits) > 0 {
						if res, _ := strconv.Atoi(api.Limits["X-Ratelimit-Remaining"][0]); res > 1 {
							if res < 50 {
								logrus.WithField("X-Ratelimit-Remaining", res).Warnln("limit less than 50")
							}
							go lookupSHA1(api, event.Name, options)
						} else {
							resetIn := api.Limits["X-RateLimit-Reset-In"][0]
							sz := len(resetIn)
							if sz > 0 && resetIn[sz-1] == 's' {
								resetIn = resetIn[:sz-1]
								sleepTime, _ := strconv.Atoi(resetIn)
								logrus.WithField("X-RateLimit-Reset-In", resetIn).Warnln("limit reached, will delay until it is reset")
								time.Sleep(time.Duration(sleepTime) * time.Second)
							}
						}
					} else {
						lookupSHA1(api, event.Name, options)
					}
				}
			case err := <-watcher.Errors:
				logrus.Fatalln(err)
			}
		}
	}()
	<-done
}

func lookupSHA1(api *api.API, filePath string, options ScanOptions) {
	file, err := os.Open(filePath)
	if err != nil {
		logrus.Warningln("Failed to read file")
		return
	}
	stat, err := file.Stat()
	defer file.Close()
	if err == nil {
		if err := filepath.Walk(filePath, addPaths); err != nil {
			logrus.Fatalln(err)
		}
	}
	if !stat.IsDir() {
		resSha1, err := utils.GetFileSHA1(filePath)
		if err == nil {
			if res, err := api.FindOrScan(filePath, resSha1, options.Headers, options.LookupFile, options.Poll); err == nil {
				logrus.Println(res)
			}
		} else {
			logrus.Println(api.ScanFile(filePath, options.Headers, options.Poll))
		}
	}
}

func addPaths(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() && !strings.Contains(path, "/.") {
		return watcher.Add(path)
	}
	return nil
}
