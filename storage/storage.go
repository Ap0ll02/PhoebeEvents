package storage

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type MyEvents struct {
	EventIDs []int
	Events   []string
}

func LoadEvents(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Couldn't Read The File: %s", filename)
		os.Exit(1)
	}
	var evnt MyEvents
	if err := toml.Unmarshal(data, &evnt); err != nil {
		fmt.Printf("Couldn't Unmarshal The Text")
		os.Exit(1)
	}
	err = CheckErr(err)

	return evnt.Events, err
}

func SaveEvents(filename string, week_day string, event string) {

}

func AddEvents(filename string, week_day string, event string) {

}

func DeleteEvent(filename string, week_day string, event string) {

}

func ClearEvents(filename string) {

}

func InitializeEvents(filename string) {
	init_data := `
[Monday]

[Tuesday]

[Wednesday]

[Thursday]

[Friday]

[Saturday]

[Sunday]
	`

	_, err := os.Stat(filename)
	if !os.IsExist(err) {
		_, err = os.Create(filename)
	}
	if err != nil {
		fmt.Printf("Could not stat file %s", filename)
		os.Exit(1)
	}
	// Save init data to initialize weekday categories in toml
	_, err = os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	if f, err := os.Stat(filename); f.Size() == 0 {
		err = os.WriteFile(filename, []byte(init_data), 0644)
		if err != nil {
			fmt.Printf("Bro it is not writing correctly")
			os.Exit(1)
		}

	} else if err != nil {
		fmt.Printf("couldnt stat file, even after checking.")
		os.Exit(1)
	}
	return
}

func CheckErr(err error) error {
	if err != nil {
		fmt.Printf("Error with the TOML file.")
		return err
	}
	return nil
}
