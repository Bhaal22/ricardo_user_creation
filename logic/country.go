package logic

import (
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/Bhaal22/ricardo_user_creation/utils"
)

var (
    validCountries = map[string]bool{"CH": true}
    Client utils.HTTPClient
)

func init() {
    Client = &http.Client{}
}

func Country(ip string) (string, error) {
    url := fmt.Sprintf("https://ipapi.co/%s/country", ip)
    request, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        return "", err
    }
    response, err := Client.Do(request)

    if err != nil {
        return "", errors.New("unable to check ip adress")
    } else {
        if response.StatusCode != 200 {
            return "", errors.New("resource not found")
        }

        defer response.Body.Close()

        body, _ := ioutil.ReadAll(response.Body)

        fmt.Printf("country %s \n", string(body))
        return string(body), nil
    }
}

func ValidCountry(country string) bool {
    return validCountries[country]
}
