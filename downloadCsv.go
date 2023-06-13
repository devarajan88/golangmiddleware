package main

import (
	"fmt"
	"io"
	util "maccsv/etc"
	"net/http"
	"os"
	"sync"
	"time"
)

func autoDownloadCSV() {

	downloadTime := time.Now().Add(5 * time.Second)

	var wg sync.WaitGroup

	multiServer := util.ReadMultiServerConfig()

	for serverNumber, url := range *multiServer {

		wg.Add(1)
		go func(u string, serverNumber int) {
			defer wg.Done()

			time.Sleep(downloadTime.Sub(time.Now()))

			filename := fmt.Sprintf("downloaded_%d.csv", time.Now().UnixNano())

			err := DownloadCSV(u, filename)
			if err != nil {
				fmt.Printf("Error downloading %s: %v\n", u, err)
			}
		}(url, serverNumber)
	}

	wg.Wait()

}

// CSV file download and saving on localstorage
func DownloadCSV(url, filename string) error {
	// Send GET request to the server
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
	return nil
}
