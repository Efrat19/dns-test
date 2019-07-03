package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)


func main() {
	n := flag.String("n", "", "lookup server")
	p := flag.Duration("p", time.Millisecond*10, "pause between lookups")

	flag.Parse()

	if *n == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	for {
		_, err := net.LookupHost(*n)
		if err != nil {
			fmt.Println("error:", err)
		}

		time.Sleep(*p)
	}
}