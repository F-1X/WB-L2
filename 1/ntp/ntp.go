package ntp

import (
	"fmt"
	"os"

	ntp "github.com/beevik/ntp"
)

func GetCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

	}
	fmt.Println(time)
}
