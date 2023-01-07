package sendMessage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	twlioClient *twilio.RestClient
)

func CreateClient() {
	twlioClient = twilio.NewRestClient()
}

func Birthday(recieverNumber string, message string) {

	senderNumber := os.Getenv("SENDER_NUMBER")

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(recieverNumber)
	params.SetFrom(senderNumber)
	params.SetBody(message)

	fmt.Println("Sending message for: " + recieverNumber)

	resp, err := twlioClient.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err)
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

}
