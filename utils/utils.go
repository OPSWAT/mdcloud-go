package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
)

// StringInSlice checks for string in slice
func StringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

// StringSlice makes a slice str to str pointers
func StringSlice(strs []string) []*string {
	res := make([]*string, len(strs))
	for i := 0; i < len(strs); i++ {
		res[i] = &(strs[i])
	}
	return res
}

// IsLetter checks if string is letter
func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// GetFileSHA1 returns a files SHA1
func GetFileSHA1(filePath string) (string, error) {
	var resSha1 string
	file, err := os.Open(filePath)
	if err != nil {
		return "", errors.Wrap(err, "Failed to open file")
	}
	defer file.Close()
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	hashInBytes := hash.Sum(nil)[:20]
	resSha1 = hex.EncodeToString(hashInBytes)
	return resSha1, nil
}

// VerifyArgsOrRun will check args length then call first defined function or 2nd to handle error
func VerifyArgsOrRun(args []string, equalTo int, call ...func()) {
	if len(args) > 0 || equalTo > 0 && len(args) == equalTo {
		call[0]()
	} else {
		if call[1] != nil {
			call[1]()
		} else {
			logger.Fatalln("args count not valid")
		}
	}
}

// FilterMap items from array based on test functions
func FilterMap(ms map[string][]string, test func(string) bool) (ret map[string][]string) {
	ret = make(map[string][]string)
	for s, ss := range ms {
		if test(s) {
			ret[s] = ss
		}
	}
	return
}
