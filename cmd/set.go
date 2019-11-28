package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/p886/byodb/commandparser"
	"github.com/p886/byodb/storage"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setCmd)
}

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set data",
	Long:  `Persist value associated with key to the database`,
	Args:  cobra.ExactValidArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Error loading .env: '%s'\n", err.Error())
		}
		storageFilePath := os.Getenv("STORAGE_FILE_PATH")
		key := args[0]
		value := args[1]
		storage.Store(storageFilePath, commandparser.Query{Command: "SET", Key: key, Value: value})
		if err != nil {
			fmt.Printf("Error getting data with key '%s': %s\n", key, err.Error())
			os.Exit(1)
		}
		fmt.Printf("Successfully saved key: '%s', value: '%s'\n", key, value)
	},
}
