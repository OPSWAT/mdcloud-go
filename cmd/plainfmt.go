package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type PlainFmt struct {
}

func (f *PlainFmt) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(fmt.Sprintf("%s\n", entry.Message)), nil
}
