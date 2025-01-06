package storage

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type WeekEvents map[string]map[int]string

func LoadEvents(filename string) (WeekEvents, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return WeekEvents{}, nil
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var events WeekEvents
	if err := toml.Unmarshal(data, &events); err != nil {
		return nil, err
	}

	if events == nil {
		events = make(WeekEvents)
	}

	for day := range events {
		if events[day] == nil {
			events[day] = make(map[int]string)
		}
	}

	return events, nil
}

func SaveEvents(filename string, events WeekEvents) error {
	data, err := toml.Marshal(events)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func AddEvent(filename, weekday, event string) error {
	events, err := LoadEvents(filename)
	if err != nil {
		return err
	}

	if _, exists := events[weekday]; !exists {
		events[weekday] = make(map[int]string)
	}

	eventID := len(events[weekday]) + 1
	events[weekday][eventID] = event
	return SaveEvents(filename, events)
}

func RemEvent(filename, weekday, event string) error {
	week_events, err := LoadEvents(filename)
	if err != nil {
		return err
	}
	day_events, exists := week_events[weekday]
	if !exists {
		return fmt.Errorf("Couldn't Find Weekday: %s", weekday)
	}
	found := false
	for id, e := range day_events {
		if e == event {
			delete(day_events, id)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("Couldn't find %s, on day %s", event, weekday)
	}

	newDayEvents := make(map[int]string)
	newID := 1
	for _, e := range day_events {
		newDayEvents[newID] = e
		newID++
	}
	week_events[weekday] = newDayEvents

	if err := SaveEvents(filename, week_events); err != nil {
		return fmt.Errorf("Failed to save events: %v", err)
	}
	fmt.Printf("Successfully Removed %s from day %s\n", event, weekday)
	return nil
}

func GetEvents(filename, weekday string) ([]string, error) {
	events, err := LoadEvents(filename)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, event := range events[weekday] {
		result = append(result, event)
	}

	return result, nil
}
