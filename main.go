package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	ttpUrl    = "https://ttp.cbp.dhs.gov/schedulerapi/slots/asLocations?minimum=1&filterTimestampBy=before&timestamp=%s&serviceName=Global%20Entry"
	daysDelta = 28
	YYYYMMDD  = "2006-01-02"
)

type location struct {
	id    int
	name  string
	alert bool
}

var locations = []location{
	{
		id:    5183,
		name:  "Chicago Ohare",
		alert: true,
	},
	{
		id:    11981,
		name:  "Chicago Field Office",
		alert: true,
	},
	{
		id:    7740,
		name:  "Milwaukee Office",
		alert: true,
	},
}

func futureTime() time.Time {
	// Return the future time struct offset by daysDelta
	return time.Now().AddDate(0, 0, daysDelta)
}

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func main() {
	formattedUrl := fmt.Sprintf("https://ttp.cbp.dhs.gov/schedulerapi/slots/asLocations?minimum=1&filterTimestampBy=before&timestamp=%s&serviceName=GlobalEntry", futureTime().Format(YYYYMMDD))
	response, err := http.Get(formattedUrl)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	jsonFormattedString, err := PrettyString(string(responseData))
	fmt.Println(formattedUrl)
	fmt.Println(jsonFormattedString)
}
