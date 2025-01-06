/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "A small CLI tool for managing events and ideas",
	Long: `
This is the Phoebe CLI Tool for event and idea management.
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

func init() {
	rootCmd.AddCommand(rootCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
