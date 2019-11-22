package storage

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/p886/byo-database/commandparser"
)

func TestStore(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "storage_test")
	if err != nil {
		t.Fatalf("Error creating tempfile: '%s'\n", err.Error())
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	cases := []struct {
		name                 string
		filePath             string
		inputQuery           commandparser.Query
		expectedError        error
		expectedFileContents []byte
	}{
		{
			name:                 "success",
			filePath:             tmpFile.Name(),
			inputQuery:           commandparser.Query{Command: "PUT", Key: "foo", Value: "bar"},
			expectedFileContents: []byte("foo bar\n"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Store(tmpFile.Name(), c.inputQuery)

			if !reflect.DeepEqual(err, c.expectedError) {
				t.Errorf("Expected error: %#v, got: %#v", c.expectedError, err)
			}

			fileContents, err := ioutil.ReadFile(tmpFile.Name())
			if err != nil {
				t.Fatalf("Error reading tempfile: '%s'\n", err.Error())
			}
			if !bytes.Equal(fileContents, c.expectedFileContents) {
				t.Errorf("Expected error: %#v, got: %#v", c.expectedFileContents, fileContents)
			}
			err = ioutil.WriteFile(tmpFile.Name(), []byte(""), 0644)
			if err != nil {
				t.Fatalf("Error truncating tempfile: '%s'\n", err.Error())
			}
		})
	}
}
