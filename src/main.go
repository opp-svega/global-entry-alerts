package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		name:  "Milwaukee",
		alert: true,
	},
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

func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "    "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}

func main() {

	for _, element := range locations {

		formattedUrl := fmt.Sprintf("https://ttp.cbp.dhs.gov/schedulerapi/slots?orderBy=soonest&limit=1&locationId=%d&minimum=1", element.id)
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
		fmt.Println(jsonFormattedString)
	}
}
