package commandparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseCommand(t *testing.T) {
	cases := []struct {
		name          string
		input         string
		expectedQuery Query
		expectedError error
	}{
		{
			name:          "General parsing error",
			input:         "FAIL",
			expectedError: fmt.Errorf("Error parsing query: '%s'", "FAIL"),
		},
		{
			name:          "PUT wrong query format 1",
			input:         "PUT asdf",
			expectedError: fmt.Errorf("Wrong query format: '%s'", "PUT asdf"),
		},
		{
			name:          "PUT wrong query format 2",
			input:         "PUT",
			expectedError: fmt.Errorf("Wrong query format: '%s'", "PUT"),
		},
		{
			name:          "PUT success",
			input:         "PUT test 123",
			expectedQuery: Query{Command: "PUT", Key: "test", Value: "123"},
		},
		{
			name:          "GET wrong query format",
			input:         "GET",
			expectedError: fmt.Errorf("Wrong query format: '%s'", "GET"),
		},
		{
			name:          "GET success",
			input:         "GET test",
			expectedQuery: Query{Command: "GET", Key: "test"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			query, err := ParseCommand(c.input)
			if !reflect.DeepEqual(query, c.expectedQuery) {
				t.Errorf("Expected query: %#v, got: %#v", c.expectedQuery, query)
			}
			if !reflect.DeepEqual(err, c.expectedError) {
				t.Errorf("Expected error: %s, got: %s", c.expectedError.Error(), err.Error())
			}
		})
	}
}
