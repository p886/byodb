package retrieval

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestRetrieve(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "storage_test")
	if err != nil {
		t.Fatalf("Error creating tempfile: '%s'\n", err.Error())
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	cases := []struct {
		Name           string
		queryKey       string
		fileContents   []byte
		expectedResult []byte
		expectedError  error
	}{
		{
			Name:           "success",
			fileContents:   []byte("foo bar"),
			queryKey:       "foo",
			expectedResult: []byte("bar"),
		},
		{
			Name:          "not found",
			fileContents:  []byte("other bar"),
			queryKey:      "foo",
			expectedError: ErrNoResult,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			// prepare tmpFile. Will be truncated on the next run, no need for clean up
			ioutil.WriteFile(tmpFile.Name(), c.fileContents, 0644)

			result, err := Retrieve(tmpFile.Name(), c.queryKey)
			if !reflect.DeepEqual(err, c.expectedError) {
				t.Errorf("Expected error: %s, got: %s", c.expectedError.Error(), err.Error())
			}
			if !reflect.DeepEqual(result, c.expectedResult) {
				t.Errorf("Expected result: %#v, got: %#v", c.expectedResult, result)
			}

		})
	}
}
