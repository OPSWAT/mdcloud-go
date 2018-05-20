package filescan

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

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
			watchDirScan(api, path[0], headers)
		case mode.IsRegular():
			fmt.Printf(api.ScanFile(path[0], headers))
		}
	}
}

func watchDirScan(api api.API, path string, headers []string) {
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(path, watchDir); err != nil {
		log.Fatalln("Error adding to watcher: ", err)
	}

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if (event.Op == fsnotify.Write) || (event.Op == fsnotify.Create) {
					log.Println(event.Op, event.Name)
					resSha1, err := getFileSha1(event.Name)
					if err != nil {
						api.ScanFile(event.Name, headers)
					} else {
						fmt.Println(api.FindOrScan(event.Name, resSha1, headers))
					}
				}
			case err := <-watcher.Errors:
				log.Fatal(err)
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
