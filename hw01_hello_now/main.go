package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	time := time.Now()
	ntptime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalf("Error: %s", err)
	}

	fmt.Printf("current time: %s\nexact time: %s\n", time, ntptime)
}
