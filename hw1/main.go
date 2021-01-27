package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
)

const (
	HOST string = "ntp1.stratum2.ru"
)

func main() {

	l := log.New(os.Stderr, "", 1)
	currentTime, err := ntp.Time(HOST)
	if err != nil {
		l.Println("error")
		os.Exit(1)
	}
	fmt.Println(currentTime)
}
