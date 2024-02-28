package main

import (
	"log"
	"time"
)

func main() {
	cc := setupClientConfig()

	log.Printf("Starting client with config: %+v", cc)
	c, err := setupClient(cc)

	if err != nil {
		log.Fatalf("Error setting up client: %v", err)
	}

	for {
		s, err := doRequest(&c, cc)

		if err != nil {
			log.Printf("Error making request: %v", err)
		} else {
			log.Printf("Request to %s returned status: %s", cc.serverURL, s)
		}
		time.Sleep(time.Duration(cc.RequestIntervalSeconds) * time.Second)
	}
}
