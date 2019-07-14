package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
	"github.com/heatxsink/go-logstash"
)

func main() {
	fmt.Println("start")
	lookupServer, duration := getFlags()
	fmt.Println(lookupServer)
	logstashServer := initLogstash()
	fmt.Println("on")
	fmt.Println("pinging ",lookupServer)
	lookup(logstashServer, lookupServer, duration)
}

func getFlags() (string, time.Duration) {
	n := flag.String("n", "", "lookup server")
	p := flag.Duration("p", time.Millisecond*10, "pause between lookups")

	flag.Parse()

	if *n == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return *n, *p
}

func initLogstash() *logstash.Logstash {
	time.Sleep(60 * time.Second)
	port, _ := strconv.Atoi(os.Getenv("LOGSTASH_PORT"))
	log := logstash.New(os.Getenv("LOGSTASH_HOST"), port, 5)
	_, err := log.Connect()
	if err != nil {
		fmt.Println(err)
	}
	logErr := log.Writeln("test")
	if logErr != nil {
		fmt.Println(logErr.Error())
	}
	return log
}

func lookup (log *logstash.Logstash, server string, duration time.Duration) {
	for {
		_, dnsErr := net.LookupHost(server)
		if dnsErr != nil {
			fmt.Println(dnsErr)
			err := log.Writeln(dnsErr.Error())
			if err != nil {
				fmt.Println("error on writing to logstash")
				fmt.Println(err)
				os.Exit(1)
			}
		}
		time.Sleep(duration)
	}
}