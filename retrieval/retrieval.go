package retrieval

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

// TODO: add mutex around file write

// ErrNoResult is returned when no result could be found
var ErrNoResult = errors.New("no result")

// Retrieve gets the data associated with a key
func Retrieve(filename string, queryKey string) ([]byte, error) {
	contents, err := ioutil.ReadFile(fmt.Sprintf("./%s", filename))
	if err != nil {
		return []byte{}, err
	}
	lines := bytes.Split(contents, []byte("\n"))

	var match []byte
	resultErr := ErrNoResult
	// Iterate through all lines, assign all matching line's value to `match`
	// Then return the last match. The last match holds the most recent write.
	for _, line := range lines {
		splitLine := bytes.SplitAfterN(line, []byte(" "), 2)
		if len(splitLine) < 2 {
			continue
		}
		key := bytes.TrimSpace(splitLine[0])
		value := bytes.TrimSpace(splitLine[1])
		if bytes.Equal(key, []byte(queryKey)) {
			match = []byte(value)
			resultErr = nil
		}
	}
	return match, resultErr
}
