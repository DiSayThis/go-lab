package main

import (
	"flag"
	"fmt"

	"github.com/DiSayThis/go-lab.git/cmd/geo"
	"github.com/DiSayThis/go-lab.git/cmd/weather"
)

func main() {
	city := flag.String("city", "Moscow", "Город")
	format := flag.Int("format", 1, "Формат вывода")
	flag.Parse()
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(geoData.City)
	weatherData, err := weather.GetWeather(*geoData, *format)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(weatherData)
}
