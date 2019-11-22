package commandparser

import (
	"fmt"
	"strings"
)

// Query is the parsed query based on user input
type Query struct {
	Command string
	Key     string
	Value   string
}

// ParseCommand parses commands given by the user
func ParseCommand(input string) (Query, error) {
	cmdParts := strings.SplitAfterN(input, " ", 3)
	for i, part := range cmdParts {
		cmdParts[i] = strings.TrimSpace(part)
	}
	if cmdParts[0] == "PUT" {
		if len(cmdParts) != 3 {
			return Query{}, fmt.Errorf("Wrong query format: '%s'", input)
		}
		return Query{
			Command: "PUT",
			Key:     cmdParts[1],
			Value:   cmdParts[2],
		}, nil
	} else if cmdParts[0] == "GET" {
		if len(cmdParts) != 2 {
			return Query{}, fmt.Errorf("Wrong query format: '%s'", input)
		}
		return Query{
			Command: "GET",
			Key:     cmdParts[1],
		}, nil
	} else {
		return Query{}, fmt.Errorf("Error parsing query: '%s'", input)
	}
}
