package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/DiSayThis/go-lab.git/cmd/geo"
)

func GetWeather(geo geo.GeoData, format int) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("Failed to parse URL")
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", errors.New("Not200")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
