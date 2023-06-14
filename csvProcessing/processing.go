package csvprocessing

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"
)

func getNewPhonesRegistered(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return 0
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	columnNames, err := reader.Read()
	if err != nil {
		fmt.Printf("Error reading column names: %v\n", err)
		return 0
	}

	registeredColumnIndex := -1
	for i, columnName := range columnNames {
		if columnName == "Registered" {
			registeredColumnIndex = i
			break
		}
	}

	if registeredColumnIndex == -1 {
		fmt.Println("Error: 'Registered' column not found")
		return 0
	}

	newPhonesRegistered := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error reading row: %v\n", err)
			break
		}

		if len(row) > registeredColumnIndex && row[registeredColumnIndex] != "" {
			registrationTime, err := time.Parse("2006-01-02", row[registeredColumnIndex])
			if err != nil {
				fmt.Printf("Error parsing registration date: %v\n", err)
				continue
			}

			if registrationTime.After(lastDownloadTime) {
				newPhonesRegistered++
			}
		}
	}

	return newPhonesRegistered
}
