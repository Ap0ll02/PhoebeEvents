package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func LoadEvents(filename string) [][]string {
	var events [][]string
	events = append(events, GetEvents(filename, "Monday"))
	events = append(events, GetEvents(filename, "Tuesday"))
	events = append(events, GetEvents(filename, "Wednesday"))
	events = append(events, GetEvents(filename, "Thursday"))
	events = append(events, GetEvents(filename, "Friday"))
	events = append(events, GetEvents(filename, "Saturday"))
	events = append(events, GetEvents(filename, "Sunday"))
	return events
}

func GetEvents(filename string, week_day string) []string {
	db, err := sql.Open("sqlite3", filename)
	CheckErr(err)

	raw_events, err := db.Query("SELECT " + week_day + " FROM events")
	CheckErr(err)

	var events []string
	var event string
	defer raw_events.Close()
	for raw_events.Next() {
		err := raw_events.Scan(&event)
		CheckErr(err)
		events = append(events, event)
		events = append(events, "\n")
	}
	err = raw_events.Err()
	CheckErr(err)
	return events
}

// func SaveEvents(filename string, week_day string, event string) {

// }

func AddEvents(filename string, week_day string, event string) {
	db, err := sql.Open("sqlite3", filename)
	CheckErr(err)
	_, err = db.Exec("INSERT INTO events (" + week_day + ") VALUES ('" + event + "')")
	CheckErr(err)
}

func DeleteEvent(filename string, week_day string, event string) {
	db, err := sql.Open("sqlite3", filename)
	CheckErr(err)
	_, err = db.Exec("DELETE FROM events WHERE " + week_day + "= '" + event + "'")
	CheckErr(err)
}

func ClearEvents(filename string) {
	db, err := sql.Open("sqlite3", filename)
	CheckErr(err)
	_, err = db.Exec("DELETE FROM events")
}

func InitializeEvents(filename string) {
	db, err := sql.Open("sqlite3", filename)
	CheckErr(err)

	// Create table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS events (Monday TEXT(65535), Tuesday TEXT(65535), Wednesay TEXT(65535), Thursday TEXT(65535), Friday TEXT(65535), Saturday TEXT(65535), Sunday TEXT(65535)")
	CheckErr(err)
}

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("Error with the DB file. %v", err)
		os.Exit(1)
	}
	return
}
