package filescan

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/OPSWAT/mdcloud-go/api"
	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

// Scan or watches files or path
func Scan(api api.API, path []string, watch bool, headers []string) {
	if path != nil && len(path) > 0 {
		fi, err := os.Stat(path[0])
		if err != nil {
			log.Fatal(err)
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			watchDir(api, path[0], headers)
		case mode.IsRegular():
			fmt.Printf(api.ScanFile(path[0], headers))
		}
	}
}

func watchDir(api api.API, path string, headers []string) {
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
					log.Println(event.Op, event.Name)
					resSha1, err := getFileSha1(event.Name)
					log.Println(resSha1)
					if err != nil {
						api.ScanFile(event.Name, headers)
					} else {
						fmt.Println(api.HashDetails(resSha1))
					}
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
