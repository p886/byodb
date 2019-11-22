package storage

import (
	"fmt"
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
	fmt.Println(tmpFile.Name())
	defer os.Remove(tmpFile.Name()) // clean up
	defer tmpFile.Close()

	cases := []struct {
		name          string
		filePath      string
		inputQuery    commandparser.Query
		expectedError error
	}{
		{
			name:       "success",
			filePath:   tmpFile.Name(),
			inputQuery: commandparser.Query{Command: "PUT", Key: "foo", Value: "bar"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := Store(tmpFile.Name(), c.inputQuery)

			if !reflect.DeepEqual(err, c.expectedError) {
				t.Errorf("Expected error: %#v, got: %#v", c.expectedError, err)
			}
		})
	}
}
