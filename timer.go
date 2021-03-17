package main

import (
	"strconv"
	"time"
)

var channelID string = ""

func timer() {
	Day := int(time.Now().Weekday())

	Today := Monday

	i := Day
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

	for {
		time.Sleep(2000)

		ActDay := int(time.Now().Weekday())
		if Day != ActDay {
			Day = ActDay

			Today = Monday

			i := Day
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

		}

		var ActTime string = strconv.FormatInt(int64(time.Now().Hour()), 10) + ":"
		buff := time.Now().Minute()

		if buff < 10 {
			ActTime += "0"
		}

		ActTime += strconv.FormatInt(int64(buff), 10)

		switch ActTime {
		case "7:59":
			AtMinute(Today[0])
		case "8:31":
			Break(Today[1], "20", "8:50")
		case "8:49":
			AtMinute(Today[1])
		case "9:21":
			Break(Today[2], "25", "9:45")
		case "9:44":
			AtMinute(Today[2])
		case "10:16":
			Break(Today[3], "5", "10:20")
		case "10:19":
			AtMinute(Today[3])
		case "10:51":
			Break(Today[4], "30", "11:20")
		case "11:19":
			AtMinute(Today[4])
		case "11:51":
			Break(Today[5], "30", "12:20")
		case "12:19":
			AtMinute(Today[5])
		case "12:51":
			Break(Today[6], "5", "12:55")
		case "12:54":
			AtMinute(Today[6])
		case "13:26":
			Break(Today[7], "5", "13:30")
		case "13:29":
			AtMinute(Today[7])
		case "14:01":
			Break(Today[8], "10", "14:10")
		case "14:09":
			AtMinute(Today[8])

		}

	}
}

func AtMinute(lesson string) {

	buff := "Za minutę " + lesson

	if lesson != "0" {
		_, _ = dg.ChannelMessageSendTTS(channelID, buff)
	}
	time.Sleep(60000)
}

func Break(lesson string, LenghtOfBreak string, TimeOfNext string) {

	buff := LenghtOfBreak + " przerwy, potem o " + TimeOfNext + " " + lesson
	if lesson != "0" {
		_, _ = dg.ChannelMessageSendTTS(channelID, buff)
	}
	time.Sleep(60000)
}
