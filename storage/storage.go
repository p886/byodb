package storage

import (
	"fmt"
	"os"

	"github.com/p886/byo-database/commandparser"
)

// Store appends the given data to the logFile
func Store(storageFilePath string, query commandparser.Query) error {
	f, err := os.OpenFile(storageFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	appendLine := fmt.Sprintf("%s %s\n", query.Key, query.Value)
	_, err = f.WriteString(appendLine)
	if err != nil {
		return err
	}
	return nil
}
