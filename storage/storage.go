package storage

import "fmt"

type MyEvents struct {
	EventIDs []int
	Day      string
	Events   []string
}

func LoadEvents(filename string) ([]string, error) {

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

}
