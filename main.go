package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kelvins/sunrisesunset"
)

func main() {
	lat, lon := 50.2700, 30.3124
	now := time.Now()
	locationStr := flag.String("loc", fmt.Sprintf("%f,%f", lat, lon), "latitude and longitude separated by comma")

	flag.Parse()

	location := strings.Split(*locationStr, ",")

	if len(location) == 2 {
		tempLat, errLat := strconv.ParseFloat(location[0], 64)
		tempLon, errLon := strconv.ParseFloat(location[1], 64)
		if errLat == nil && errLon == nil {
			lat = tempLat
			lon = tempLon
		}
	}

	_, offset := now.Zone()
	p := sunrisesunset.Parameters{
		Latitude:  lat,
		Longitude: lon,
		Date:      now,
		UtcOffset: float64(offset / 60 / 60),
	}

	// Calculate the sunrise and sunset times
	sunrise, sunset, err := p.GetSunriseSunset()

	// If no error has occurred, print the results
	if err == nil && now.Unix() < sunrise.Unix() || now.Unix() > sunset.Unix() {
		fmt.Println("Night")
	} else {
		fmt.Println("Day")
	}
}
