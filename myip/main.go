package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

func init() {
	var err error
	db, err = geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer db.Close()
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/text", []byte(ctx.ClientIP()+"\n"))
	})

	r.GET("/geo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, lookup(ctx.ClientIP()))
	})

	r.Run(":8080")
}

func lookup(addr string) any {
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
