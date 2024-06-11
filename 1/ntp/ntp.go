package ntp

import (
	ntp "github.com/beevik/ntp"
)

func GetCurrentTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		panic(err)
	}
	println(time)
}
