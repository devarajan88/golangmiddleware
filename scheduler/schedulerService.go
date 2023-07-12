package scheduler

import (
	"log"
	"time"
)

func runJob() {

	ticker := time.NewTicker(24 * time.Hour)

	tickerChan := ticker.C

	go func() {
		for {
			<-tickerChan

			//TODO: Implement the goal of the job
			log.Println("Scheduled Job executed at ", time.Now())
		}
	}()

	// Sleep forever to keep the main goroutine alive
	select {}
}
