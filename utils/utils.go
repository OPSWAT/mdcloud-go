package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"github.com/pkg/errors"
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

// GetFileSHA1 returns file sha1
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
