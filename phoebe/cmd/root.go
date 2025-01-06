/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"phoebe/storage"
)

var tomlPath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "phoebe",
	Aliases: []string{"pb"},
	Version: "1.0",
	Short:   `Phoebe is a terminal agenda CLI to manage your events`,
	Long: `
Phoebe CLI: A Simple Tool For Mantaining A Weekly Agenda.
		
	Usage:
		pb [command] [arguments] [flags]
	
	Available Commands:
		add    Add a new event to a specific day
		week   View the full week, or specified day
		del    Remove an event from a specified day
		help   Help about a command
	
	Flags:
		-h, --help   Show help for the CLI
		`,
	Example: `
	# Add an event to Monday:
	pb add Monday "Meeting at the cafe"

	# View all the events for the week
	pb week

	# View all the events for Monday
	pb week monday

	# Remove an event
	pb del Monday "Meeting at the cafe"
		`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		err := storage.InitializeEventsFile(tomlPath)
		if err != nil {
			fmt.Println("Couldn't Initialize Events File")
			os.Exit(1)
		}
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.phoebe.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVar(&tomlPath, "toml", "events.toml", "Path To The Events TOML File")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(completionCmd)
}

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish]",
	Short: "Generate completion script",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Specify shell type: bash, zsh, or fish")
			return
		}

		switch args[0] {
		case "bash":
			rootCmd.GenBashCompletion(os.Stdout)
		case "zsh":
			rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			rootCmd.GenFishCompletion(os.Stdout, true)
		default:
			fmt.Printf("Unsupported Shell Type: %s\n", args[0])
		}
	},
}
