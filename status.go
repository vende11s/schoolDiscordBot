//Yes it's bad code but I dunno go
package main

import (
	"time"
)

var Timers [18]time.Time

var Lessons []string = []string{"8h", "8h30m", "8h50m", "9h20m", "9h45m", "10h15m", "10h20m", "10h50m", "11h20m", "11h50m", "12h20m", "12h50m", "12h55m", "13h25m", "13h30m", "14h", "14h10m", "14h40m"}

func status() {

	j := time.Now()
	j = j.Add(time.Hour * time.Duration(time.Now().Hour()) * -1)
	j = j.Add(time.Minute * time.Duration(time.Now().Minute()) * -1)
	j = j.Add(time.Second * time.Duration(time.Now().Second()) * -1)

	for i := 0; i < 18; i++ {
		Timers[i] = j
		jd, _ := time.ParseDuration(Lessons[i])
		Timers[i] = Timers[i].Add(jd)
	}

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

	i = 0
	slip := true
	for {
		if slip {
			time.Sleep(time.Second * 5)
		}
		do := time.Until(Timers[i])
		if do < 0 {
			i++
			slip = false
		} else {
			slip = true
			do = do.Round(time.Second)
			if i%2 == 0 {
				dg.UpdateListeningStatus(do.String() + " do " + Today[i/2])
			} else {
				dg.UpdateListeningStatus(do.String() + " do Przerwy")
			}
		}
		if int(time.Now().Weekday()) != Day {
			Day = int(time.Now().Weekday())
			switch Day {
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
	}

}
