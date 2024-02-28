package main

import "log"

func main() {

	sc := setupServerConfig()

	log.Printf("Starting server with config: %+v", sc)
	log.Fatal(runServer(sc))
}
