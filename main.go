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
	for _, birthday := range activeBirthdays {
		randomMessage := messageGenerator.GenerateMessage(birthday.Name)
		println(randomMessage)
		// sendMessage.Birthday(birthday.Phone, randomMessage)
	}
}

func scheduler() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Day().At("10:30").Do(func() {
		alertBirthday(checkDate(birthdayLoad()))
	})
	s.StartBlocking()
}

func main() {
	var enviroment string
	flag.StringVar(&enviroment, "enviroment", "production", "Whether in proudction or development. Default: production")

	flag.Parse()

	switch enviroment {
	case "development":
		envLoad()
	case "production":

	default:
		log.Fatal("Unknown enviroment!")
	}

	sendMessage.CreateClient()

	scheduler()

}
