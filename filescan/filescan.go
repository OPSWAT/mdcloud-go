package filescan

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/OPSWAT/mdcloud-go/api"
	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

// Scan or watches files or path
func Scan(api api.API, path []string, watch bool, headers []string) {
	if path != nil && len(path) > 0 {
		fi, err := os.Stat(path[0])
		if err != nil {
			logrus.Fatalln(err)
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			watchDirScan(api, path[0], headers)
		case mode.IsRegular():
			if res, err := api.ScanFile(path[0], headers); err == nil {
				logrus.Println(res)
			} else {
				logrus.Fatalln(err)
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
					resSha1, err := getFileSha1(event.Name)
					if err != nil {
						api.ScanFile(event.Name, headers)
					} else {
						if res, err := api.FindOrScan(event.Name, resSha1, headers); err == nil {
							logrus.Println(res)
						} else {
							logrus.Fatalln(err)
						}
					}
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

func getFileSha1(filePath string) (string, error) {
	var resSha1 string
	file, err := os.Open(filePath)
	if err != nil {
		return resSha1, err
	}
	defer file.Close()
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return resSha1, err
	}
	hashInBytes := hash.Sum(nil)[:20]
	resSha1 = hex.EncodeToString(hashInBytes)
	return resSha1, nil
}
