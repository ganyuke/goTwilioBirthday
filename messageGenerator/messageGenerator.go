package messageGenerator

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type PreMessage struct {
	Salutations  []string `json:"salutations"`
	Body         []string `json:"body"`
	Valedictions []string `json:"valedictions"`
	SenderName   []string `json:"sender_name"`
}

func loadPreMessages() *PreMessage {
	pathToMessages := os.Getenv("MESSAGE_TEMPLATE_PATH")
	messageGenFile, err := os.ReadFile(pathToMessages)
	if err != nil {
		log.Fatal(err)
	}

	var messages *PreMessage
	err = json.Unmarshal(messageGenFile, &messages)
	if err != nil {
		log.Fatal(err)
	}

	return messages
}

func randomizer(length int) int {
	if length <= 1 {
		return 0
	}

	rand.Seed(time.Now().UnixNano())

	return rand.Intn(length)

}

func GenerateMessage(name string) string {
	messages := loadPreMessages()

	salutations := messages.Salutations
	salutation := salutations[randomizer(len(salutations))]

	bodys := messages.Body
	body := bodys[randomizer(len(bodys))]

	valedictions := messages.Valedictions
	valediction := valedictions[randomizer(len(valedictions))]

	senderNames := messages.SenderName
	senderName := senderNames[randomizer(len(senderNames))]

	message := salutation + body + valediction + senderName

	message = strings.Replace(message, "{name}", name, -1)

	return message
}
