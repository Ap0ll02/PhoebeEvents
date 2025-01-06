/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"phoebe/storage"
	"strings"

	"github.com/spf13/cobra"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Short: "Delete a specified event from a given day.",
	Long: `
Delete an event that you specify, from a specified day of the week.

Usage:
	pb del [weekday] ["event name"]
		`,
	Example: `
# Remove a coffee shop date event from Friday
pb del Friday "Coffee Shop Date"
		`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		event := args[1]
		day := args[0]
		day = strings.ToUpper(day)
		err := storage.RemEvent(tomlPath, day, event)
		if err != nil {
			fmt.Printf("There was an error with your request: %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
