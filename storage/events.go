package storage

import (
	"fmt"
	"os"
)

func InitializeEventsFile(filePath string) error {
	_, err := os.Stat(filePath)
	if err == nil {
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("Could not open TOML File: %v", err)
		}
		defer file.Close()

		info, err := file.Stat()
		if err != nil {
			return fmt.Errorf("Couldn't Retrieve File Info: %v", err)
		}
		if info.Size() == 0 {
			return WriteDefaultFile(filePath)
		}
	} else if os.IsNotExist(err) {
		return WriteDefaultFile(filePath)
	} else {
		return fmt.Errorf("Error Cheking File: %v", err)
	}
	return nil
}

func WriteDefaultFile(filePath string) error {
	default_data := `
[MONDAY]

[TUESDAY]

[WEDNESDAY]

[THURSDAY]

[FRIDAY]

[SATURDAY]

[SUNDAY]
`
	err := os.WriteFile(filePath, []byte(default_data), 0644)
	if err != nil {
		return fmt.Errorf("could not create toml file: %v", err)
	}
	fmt.Printf("Wrote To TOML File %s and initialized", filePath)
	return nil
}
