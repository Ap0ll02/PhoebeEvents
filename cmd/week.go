/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"phoebe/storage"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// weekCmd represents the week command
var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Print out the schedule for the week, or specify a day.",
	Long: `
# Sets up the entire weekly agenda for a given week:

Usage:
	pb week [weekday (optional)]
		`,
	Example: `
# View the events for Friday
pb week Friday

# View the events of the week
pb week
		`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		d := color.New(color.FgCyan, color.Bold)
		b := color.New(color.FgGreen, color.Bold)
		c := color.New(color.FgRed, color.Bold)
		e := color.RGB(250, 90, 250).Add(color.Bold)
		f := color.New(color.FgYellow, color.Bold)
		g := color.RGB(250, 100, 170).Add(color.Bold)
		h := color.New(color.FgBlue, color.Bold)

		color_list := []color.Color{*b, *c, *d, *e, *f, *g, *h}
		week_days := []string{"MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"}
		if len(args) == 0 {
			w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', 0)
			// fmt.Fprintln(w, "MONDAY\tTUESDAY\tWEDNESDAY\tTHURSDAY\tFRIDAY\tSATURDAY\tSUNDAY\t")
			events, err := storage.LoadEvents(tomlPath)
			if err != nil {
				os.Exit(1)
			}
			for ind, day := range week_days {
				day_events, exists := events[day]
				event_list := "No Events"
				if exists {
					event_list = FormatEvents(day_events)
				}
				color_list[ind].Fprintf(w, "%s: %s\t\n", day, event_list)
				// fmt.Fprintf(w, "%s:\t%s\t\n", day, event_list)
			}
			w.Flush()

		} else {
			week_day := args[0]
			week_day = strings.ToUpper(week_day)
			fmt.Printf("Schedule For %s:\n", week_day)
			day_events, err := storage.GetEvents(tomlPath, week_day)
			if err != nil {
				_ = fmt.Errorf("Couldn't Retrieve Events: %v", err)
				os.Exit(1)
			}
			for ind, event := range day_events {
				fmt.Printf("> %d: %s\n", ind, event)
			}
		}
	},
}

func FormatEvents(day_events map[int]string) string {
	if len(day_events) == 0 {
		return "No Events"
	}
	var events []string
	for _, event := range day_events {
		events = append(events, event)
	}
	return strings.Join(events, ", ")
}

func init() {
	rootCmd.AddCommand(weekCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weekCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weekCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
