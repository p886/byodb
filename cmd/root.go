package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/p886/byodb/repl"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "byodb",
	Short: "A key value database",
	Long: `A key value database I built as a learning project.
Not intended for production use.`,
	Run: func(cmd *cobra.Command, args []string) {
		envFile, err := cmd.Flags().GetString("env")
		if err != nil {
			log.Fatalf("Error getting env flag: %s", err.Error())
		}
		err = godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Error loading .env: '%s'\n", err.Error())
		}
		storageFilePath := os.Getenv("STORAGE_FILE_PATH")
		log.Printf("Using '%s' as backend", storageFilePath)
		repl.Loop(storageFilePath)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("env", "e", ".env", "Location of .env file for config")
}
