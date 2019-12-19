package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {
	response, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(response.Format(time.RFC1123Z))
}
