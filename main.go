package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"bowenchen.xyz/goTwilioBirthday/messageGenerator"
	"bowenchen.xyz/goTwilioBirthday/sendMessage"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

type Birthday struct {
	Name  string `json:"name"`
	Date  string `json:"date"`
	Phone string `json:"phone"`
}

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(".env file loaded.")
}

func sanityCheck() {
	requiredVaribles := []string{
		"TWILIO_ACCOUNT_SID",
		"TWILIO_AUTH_TOKEN",
		"SENDER_NUMBER",
		"BIRTHDAY_LIST_PATH",
		"MESSAGE_TEMPLATE_PATH",
	}

	for _, variable := range requiredVaribles {
		if os.Getenv(variable) == "" {
			log.Fatal(variable + " is not specified!")
		}
	}

}

func birthdayLoad() []Birthday {

	birthdayPath := os.Getenv("BIRTHDAY_LIST_PATH")
	birthdayFile, err := os.ReadFile(birthdayPath)
	if err != nil {
		log.Fatal("Error loading birthday list file: ", err)
	}
	birthdays := []Birthday{}
	err = json.Unmarshal(birthdayFile, &birthdays)
	if err != nil {
		log.Fatal("Error unmarshalling birthday list file: ", err)
	}

	return birthdays

}

func checkDate(birthdayList []Birthday) []Birthday {
	today := time.Now().Format("01-02")
	var activeBirthdays []Birthday
	for _, birthday := range birthdayList {
		birthDate, err := time.Parse("2006-01-02", birthday.Date)
		if err != nil {
			fmt.Println("Skipping " + birthday.Name + ". Unable to parse date: " + err.Error())
			continue
		}
		if today == birthDate.Format("01-02") {
			activeBirthdays = append(activeBirthdays, birthday)
		}
	}
	return activeBirthdays
}

func alertBirthday(activeBirthdays []Birthday) {
	fmt.Printf("\nToday is %s. There is currently %d birthday(s).\n", time.Now().Local().Format(time.RFC1123), len(activeBirthdays))
	for _, birthday := range activeBirthdays {
		randomMessage := messageGenerator.GenerateMessage(birthday.Name)
		sendMessage.Birthday(birthday.Phone, randomMessage)
	}
}

func scheduler(crontime string) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At(crontime).Do(func() {
		alertBirthday(checkDate(birthdayLoad()))
	})
	s.StartBlocking()
}

func main() {
	var enviroment string
	var twilio bool
	var crontime string

	flag.StringVar(&enviroment, "env", "prod", "If set to \"dev\", uses .env file in project root.")
	flag.BoolVar(&twilio, "text", true, "Controls whether texts will be sent or merely logged.")
	flag.StringVar(&crontime, "time", "00:00", "Time in UTC to check birthdays and send texts.")

	flag.Parse()

	switch enviroment {
	case "dev":
		envLoad()
	case "prod":
	default:
		log.Panic("Unknown enviroment!")
	}

	sanityCheck()

	if !twilio {
		fmt.Println("Message sending disabled.")
	}

	if crontime != "00:00" {
		fmt.Println("Cronjob adjusted to run at " + crontime + " UTC.")
	}

	sendMessage.CreateClient(twilio)

	scheduler(crontime)

}
