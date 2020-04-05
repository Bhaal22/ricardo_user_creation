package logic

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

var validCountries = map[string]bool{"CH": true}

func Country(ip string) (string, error) {
	fmt.Printf("https://ipapi.co/%s/country", ip)
	response, err := http.Get(fmt.Sprintf("https://ipapi.co/%s/country", ip))

	if err != nil {
		return "", errors.New("unable to check ip adress")
	} else {
		defer response.Body.Close()

		body, _ := ioutil.ReadAll(response.Body)

		return string(body), nil
	}
}

func ValidCountry(country string) bool {
	return validCountries[country]
}
