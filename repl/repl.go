package repl

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/p886/byo-database/commandparser"
	"github.com/p886/byo-database/retrieval"
	"github.com/p886/byo-database/storage"
)

// Loop continually asks for user input and executes it
func Loop(logFile string) {
	for {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		query, err := commandparser.ParseCommand(text)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Received query key: '%s', value: '%s'\n", query.Key, query.Value)
		if query.Command == "PUT" {
			storeErr := storage.Store(logFile, query)
			if err != nil {
				log.Printf("Error storing '%s'", storeErr.Error())
				continue
			}
			log.Printf("Successfully stored key: '%s', value: '%s'\n", query.Key, query.Value)
		} else if query.Command == "GET" {
			result, retrieveErr := retrieval.Retrieve(logFile, query.Key)
			if retrieveErr != nil {
				if retrieveErr == retrieval.ErrNoResult {
					log.Printf("No result for key: '%s'\n", query.Key)
					continue
				}
				log.Printf("Error retrieving key: '%s'\n", query.Key)
			}
			log.Printf("Query result: '%s'\n", result)
		}
		fmt.Println("")
	}
}
