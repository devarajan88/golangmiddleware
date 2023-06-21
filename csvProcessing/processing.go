package csvprocessing

import (
	"fmt"
	"maccsv/csv"
	"time"
)

func GetNewPhonesRegistered(iterator csv.RowIterator, lastDownloadTime time.Time) []string {
	columnNames := iterator.Get()

	registeredColumnIndex := -1
	for i, columnName := range columnNames {
		if columnName == "Registered" {
			registeredColumnIndex = i
			break
		}
	}

	if registeredColumnIndex == -1 {
		fmt.Println("Error: 'Registered' column not found")
		return nil
	}

	newPhonesRegistered := []string{}

	for iterator.Next() {
		row := iterator.Get()
		if len(row) > registeredColumnIndex && row[registeredColumnIndex] != "" {
			registrationTime, err := time.Parse("2006-01-02", row[registeredColumnIndex])
			if err != nil {
				fmt.Printf("Error parsing registration date: %v\n", err)
				continue
			}

			if registrationTime.After(lastDownloadTime) {
				newPhonesRegistered = append(newPhonesRegistered, row[registeredColumnIndex])
			}
		}
	}

	return newPhonesRegistered
}
