/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"phoebe/storage"

	"github.com/spf13/cobra"
)

var dbPath string

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A small CLI tool for managing events and ideas",
	Long: `
This is the Phoebe CLI Tool for event and idea management.
`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		storage.InitializeEvents(dbPath)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", "events.db", "The file for storing events")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
