package main

import (
	csvOriginal "encoding/csv"
	"fmt"
	"io"
	"maccsv/csv"
	csvPro "maccsv/csvprocessing"
	util "maccsv/etc"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var lastDownloadTime time.Time

func autoDownloadCSV() {

	downloadTime := time.Now().Add(5 * time.Second)
	// Currently for testing it is hard coded
	// after that we will take it from config file.

	var wg sync.WaitGroup

	multiServer := util.ReadMultiServerConfig()

	for serverNumber, url := range *multiServer {

		wg.Add(1)
		go func(u string, serverNumber int) {
			defer wg.Done()

			time.Sleep(downloadTime.Sub(time.Now()))

			currentTime := time.Now()
			formattedTime := currentTime.Format("02_01_2006")

			filename := fmt.Sprintf("downloaded_server%d_%v.csv", serverNumber, formattedTime)

			err := DownloadCSV(u, filename)
			if err != nil {
				fmt.Printf("Error downloading %s: %v\n", u, err)
			}

			csvString, err := LoadCSVAsString(filename)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			// Create a new CSV object
			csvObj, err := csv.New(csvString)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			iterator := csvObj.RowIterator(0)
			newPhones := csvPro.GetNewPhonesRegistered(iterator, lastDownloadTime)
			fmt.Println(newPhones)

		}(url, serverNumber)
	}

	wg.Wait()

}

func DownloadCSV(url, filename string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s\n", filename)

	lastDownloadTime = time.Now()
	return nil
}

func LoadCSVAsString(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File reading error")
		return "", err
	}
	defer file.Close()

	reader := csvOriginal.NewReader(file)
	reader.TrimLeadingSpace = true

	reader.LazyQuotes = true
	reader.ReuseRecord = true

	var lines []string
	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading CSV file")
			return "", err
		}

		for i := range row {
			row[i] = strings.Trim(row[i], ` "`)
		}

		line := strings.Join(row, ",")
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n"), nil
}
