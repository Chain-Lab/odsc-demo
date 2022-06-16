package utils

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path/filepath"
)

type Formatter struct{}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer

	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s - %s] %s (%s:%d)\n", entry.Level, timestamp, entry.Message, fName, entry.Caller.Line)
	} else {
		newLog = fmt.Sprintf("[%s - %s] %s\n", entry.Level, timestamp, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}
