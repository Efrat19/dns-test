package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
	"github.com/heatxsink/go-logstash"
)


func main() {
	n := flag.String("n", "", "lookup server")
	p := flag.Duration("p", time.Millisecond*10, "pause between lookups")

	flag.Parse()

	if *n == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// log := logstash.New(os.getEnv("LOGSTASH_HOST"), os.getEnv("LOGSTASH_PORT"), 5)
	log := logstash.New("172.17.0.1", 9600, 5)
	_, err := log.Connect()
	if err != nil {
		fmt.Println(err)
	}
	for {
		_, dnsErr := net.LookupHost(*n)
		err = log.Writeln(dnsErr.Error())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("ping")
		time.Sleep(*p)
	}
}
