package sendMessage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

var (
	messageService *Service
)

type Service struct {
	twlioClient *twilio.RestClient
	enabled     bool
}

func CreateClient(enabled bool) {
	messageService = &Service{twlioClient: twilio.NewRestClient(), enabled: enabled}
}

func Birthday(recieverNumber string, message string) {

	senderNumber := os.Getenv("SENDER_NUMBER")

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(recieverNumber)
	params.SetFrom(senderNumber)
	params.SetBody(message)

	if !messageService.enabled {
		fmt.Println("Message content to " + recieverNumber + " was:\n" + message)
	}

	fmt.Println("Sending message for: " + recieverNumber)

	resp, err := messageService.twlioClient.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err)
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

}
