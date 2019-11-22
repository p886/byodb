package storage

import (
	"fmt"
	"os"

	"github.com/p886/byo-database/commandparser"
)

// Store appends the given data to the logFile
func Store(fileName string, query commandparser.Query) error {
	f, err := os.OpenFile(fmt.Sprintf("./%s", fileName),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(fmt.Sprintf("%s\n", fmt.Sprintf("%s %s", query.Key, query.Value))); err != nil {
		return err
	}
	return nil
}
