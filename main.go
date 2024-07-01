package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
	"gopkg.in/yaml.v3"
)

func main() {
	db, err := geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP("24.139.142.187")
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", record.City.Names["en"])
	fmt.Printf("%v\n", record.Country.Names["en"])
	fmt.Printf("%v,%v\n", record.Location.Latitude, record.Location.Longitude)

	fmt.Println()

	res := lookup("127.0.0.1")
	y, err := yaml.Marshal(res)

	fmt.Println(string(y))
}

func lookup(addr string) any {
	db, err := geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ip := net.ParseIP(addr)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	return map[string]any{
		"ip":      addr,
		"city":    fmt.Sprintf("%v", record.City.Names["en"]),
		"country": fmt.Sprintf("%v", record.Country.Names["en"]),
		"lat":     record.Location.Latitude,
		"long":    record.Location.Longitude,
	}
}
