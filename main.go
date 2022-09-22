package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	twilio "github.com/opp-svega/global-entry-alerts/pkg/twilio"
)

type location struct {
	id         int
	name       string
	alert      bool
	outputData []data
}

type data struct {
	LocationId     int    `json:"locationId"`
	StartTimestamp string `json:"startTimestamp"`
	EndTimestamp   string `json:"endTimestamp"`
	Active         bool   `json:"active"`
	Duration       int    `json:"duration"`
	RemoteInd      bool   `json:"remoteInd"`
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
		name:  "Milwaukee",
		alert: true,
	},
	// Next Two Locations are just tests to ensure we pull back any results.
	// Anchorage Alaska has been pretty good about keeping locations open.
	{
		id:    5023,
		name:  "Detroit",
		alert: false,
	},
	{
		id:    7540,
		name:  "Anchorage Alaska",
		alert: false,
	},
}

func main() {

	log.Println("Checking all hardcoded locations...")
	for _, element := range locations {

		formattedUrl := fmt.Sprintf("https://ttp.cbp.dhs.gov/schedulerapi/slots?orderBy=soonest&limit=1&locationId=%d&minimum=1", element.id)
		response, err := http.Get(formattedUrl)

		if err != nil {
			log.Fatal(err)
		}
		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var locationData []data
		err = json.Unmarshal(responseData, &locationData)
		if err != nil {
			log.Fatal(err)
		}

		element.outputData = locationData
		if len(locationData) > 0 && element.alert == true {
			msg := fmt.Sprintf("Global Entry Appointment Found\nLocation: %s\nLocation ID: %d\nStart Time: %s", element.name, element.id, element.outputData[0].StartTimestamp)
			log.Println(msg)
			twilio.SendSMS(msg)
		}
	}
}
