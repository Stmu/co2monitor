package main

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/stmu/co2monitor/meter"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	device     = kingpin.Arg("device", "CO2 Meter device, such as /dev/hidraw2").Required().String()
	listenAddr = kingpin.Arg("listen-address", "The address to listen on for HTTP requests.").Default(":8080").String()
)

var (
	temperature = 0.0
	co2 = 0
)


func Metrics(w http.ResponseWriter, request *http.Request) {
  
    w.Write([]byte("Hello, " +  temperature + "!"))
}


func main() {
	kingpin.Parse()
	http.Handle("/metrics", Metrics)
	go measure()
	log.Printf("Serving metrics at '%v/metrics'", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}

func measure() {
	meter := new(meter.Meter)
	err := meter.Open(*device)
	if err != nil {
		log.Fatalf("Could not open '%v'", *device)
		return
	}

	for {
		result, err := meter.Read()
		if err != nil {
			log.Fatalf("Something went wrong: '%v'", err)
		}
		temperature:= result.Temperature
		co2 := float64(result.Co2)
	}
}
