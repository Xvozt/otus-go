package main

import (
	"flag"
	"fmt"
	"github.com/xvozt/otus-go/hw2/unpacker"
	"log"
)

var s string

func init() {
	flag.StringVar(&s, "string", "", "")
}

func main() {
	flag.Parse()
	if unpacked, err := unpacker.Unpack(s); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(unpacked)
	}
}
