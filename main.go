package main

import (
	"bufio"
	"fmt"
	"github.com/p886/byo-database/commandparser"
	"github.com/p886/byo-database/retrieval"
	"github.com/p886/byo-database/storage"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const logFile = "data.log"

// TODO: implement vacuuming :D

func main() {
	log.Println("Booting…")
	err := boot()
	if err != nil {
		log.Printf("error determining log file status: %s", err.Error())
		os.Exit(1)
	}
	fmt.Println()

	fmt.Println("Welcome! Enter command prefixed with PUT to store, GET to retrieve.")
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

func boot() error {
	logFileExists, err := checkLogFileExistance()
	if err != nil {
		return err
	}
	if logFileExists {
		log.Printf("Log file '%s' was found", logFile)
	} else {
		log.Printf("Log file '%s' MISSING", logFile)
		log.Println("Initializing log file…")
		err = initializeLogFile()
		if err != nil {
			return err
		}
	}
	return nil
}

func checkLogFileExistance() (bool, error) {
	if _, err := os.Stat(fmt.Sprintf("./%s", logFile)); err == nil {
		return true, nil

	} else if os.IsNotExist(err) {
		return false, nil

	} else {
		return false, err

	}
}

func initializeLogFile() error {
	filepath := fmt.Sprintf("./%s", logFile)
	return ioutil.WriteFile(filepath, []byte(""), 0644)
}
