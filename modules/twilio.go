package twilio

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

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

	resp, err := twilioConnection.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}
