/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"phoebe/storage"
	"strings"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an event to a given weekday, [weekday] [event]",
	Long: `
Add an event to a given weekday, duplicates are not allowed. Type your event in quotations.

Usage:
	pb add [weekday] ["event"]
		`,
	Example: `
# Add an event to Thursday
pb add Thursday "Compile Rust Project"
		`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		week_day := args[0]
		week_day = strings.ToUpper(week_day)
		event := args[1]
		valid_days := map[string]bool{
			"MONDAY": true, "TUESDAY": true, "WEDNESDAY": true, "THURSDAY": true, "FRIDAY": true, "SATURDAY": true, "SUNDAY": true,
		}
		if !valid_days[week_day] {
			fmt.Printf("Invalid Weekday: %s\n", week_day)
			os.Exit(1)
		}
		fmt.Println("Event Added Successfully")
		fmt.Printf("%s: %s\n", week_day, event)

		storage.AddEvent(tomlPath, week_day, event)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
