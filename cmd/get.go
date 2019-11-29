package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/p886/byodb/retrieval"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get data",
	Long:  `Get data associated with a key`,
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Error loading .env: '%s'\n", err.Error())
		}
		storageFilePath := os.Getenv("STORAGE_FILE_PATH")
		key := args[0]
		result, err := retrieval.Retrieve(storageFilePath, key)
		if err != nil {
			fmt.Printf("Error getting data with key '%s': %s\n", key, err.Error())
			os.Exit(1)
		}
		fmt.Println(string(result))
	},
}
