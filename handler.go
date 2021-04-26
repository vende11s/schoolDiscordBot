package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const token string = ""

var BotID string

var Czas string = "1. 8:00-8:30\n2. 8:50-9:20\n3. 9:45-10:15\n4. 10:20-10:50\n5. 11:20-11:50\n6. 12:20-12:50\n7. 12:55-13:25\n8. 13:30-14:00\n9. 14:10-14:40"

//Plan
var countOfLessons int = 9
var Monday []string = []string{"", "", "", "", "", "", "", "", ""}
var Tuesday []string = []string{"", "", "", "", "", "", "", "0", "0"}
var Wednesday []string = []string{"", "", "", "", "", "", "", "", ""}
var Thursday []string = []string{"0", "", "", "", "", "", "", "0", "0"}
var Friday []string = []string{"", "", "", "", "", "", "", "0", "0"}
var Saturday []string = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}
var Sunday []string = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0"}

//Days of week in Polish
var DaysOfWeek []string = []string{"Poniedziałek", "Wtorek", "Środa", "Czwartek", "Piątek", "Sobota", "Niedziela"}
var TimeOfLessons []string = []string{"8:00-8:30", "8:50-9:20", "9:45-10:15", "10:20-10:50", "11:20-11:50", "12:20-12:50", "12:55-13:25", "13:30-14:00", "14:10-14:40"}

var dg, err = discordgo.New("Bot " + token)

func main() {

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}
	BotID = u.ID

	dg.AddHandler(messageHandler)

	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(time.Now().Date())

	go timer()
	go status()

	<-make(chan struct{})

}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(m.Content)
	if m.Author.ID == BotID {
		return
		//COMMANDS
	} else if m.Content == "!czas" {
		SendEmbed(s, m.ChannelID, "", "Czasy:", Czas)
	} else if m.Content == "!plan" {

		Today := Monday

		i := int(time.Now().Weekday())
		switch i {
		case 0:
			Today = Sunday
		case 1:
			Today = Monday
		case 2:
			Today = Tuesday
		case 3:
			Today = Wednesday
		case 4:
			Today = Thursday
		case 5:
			Today = Friday
		case 6:
			Today = Saturday
		case 7:
			Today = Sunday
		}

		j := 0
		buff := ""
		for j < countOfLessons {
			buff1 := strconv.FormatInt(int64(j+1), 10)

			buff += "**" + buff1 + ". " + Today[j] + "**" + "\n"
			buff += TimeOfLessons[j] + "\n"
			j++
		}
		buff1 := "Dzisiaj jest " + DaysOfWeek[i-1]
		fmt.Println(buff)
		SendEmbed(s, m.ChannelID, "", buff1, buff)
	} else if m.Content == "!jutro" {

		Today := Monday

		i := (int(time.Now().Weekday()) + 1) % 7
		switch i {
		case 0:
			Today = Sunday
		case 1:
			Today = Monday
		case 2:
			Today = Tuesday
		case 3:
			Today = Wednesday
		case 4:
			Today = Thursday
		case 5:
			Today = Friday
		case 6:
			Today = Saturday
		case 7:
			Today = Sunday
		}

		j := 0
		buff := ""
		for j < countOfLessons {
			buff1 := strconv.FormatInt(int64(j+1), 10)

			buff += "**" + buff1 + ". " + Today[j] + "**" + "\n"
			buff += TimeOfLessons[j] + "\n"
			j++
		}
		buff1 := "Dzisiaj jest " + DaysOfWeek[i-1]
		fmt.Println(buff)
		SendEmbed(s, m.ChannelID, "", buff1, buff)
	}
}
