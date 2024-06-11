package ntp

import (
	"fmt"
	"log"

	ntp "github.com/beevik/ntp"
)

func GetCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(time)
}
