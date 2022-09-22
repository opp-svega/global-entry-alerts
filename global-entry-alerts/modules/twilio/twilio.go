//module github.com/opp-svega/global-entry-alerts/global-entry-alerts/modules/twilio

package twilio

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type response struct {
	Status      string `json:"status"`
	DateUpdated string `json:"date_updated"`
}

var twilioConnection twilio.RestClient

func init() {
	// This will look for `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` variables inside the current environment to initialize the constructor
	twilioConnection = *twilio.NewRestClient()

}

func SendSMS(message string) {
	from := os.Getenv("TWILIO_FROM_PHONE_NUMBER")
	to := os.Getenv("TWILIO_TO_PHONE_NUMBER")
	params := &openapi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(message)

	var responseData response

	resp, err := twilioConnection.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {

		response, _ := json.Marshal(*resp)
		json.Unmarshal(response, &responseData)
		logMessage := fmt.Sprintf("SMS Message is %s at %s", responseData.Status, responseData.DateUpdated)
		log.Println(logMessage)
	}
}
