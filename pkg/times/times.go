package times

import (
	"fmt"
	"time"
)

//The time package in Go provides functionality for measuring and displaying time. It supports:

//Getting the Current Time

func getCurrentTime() {
	now := time.Now()
	fmt.Println(now)

	//OutPut Current time: 2023-10-09 14:05:06.123456789 -0700 PDT m=+0.000000000

	//types
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())
	fmt.Println("Nanosecond:", now.Nanosecond())
	fmt.Println("Location:", now.Location())
	fmt.Println("Weekday:", now.Weekday())
}

// Formatting and Parsing Times
// 1- Formatting
func formatTime() {
	now := time.Now()
	fmt.Println("Default format:", now.String())
	// Custom formats
	fmt.Println("YYYY-MM-DD:", now.Format("2006-01-02"))
	fmt.Println("MM/DD/YYYY:", now.Format("01/02/2006"))
	fmt.Println("HH:MM:SS:", now.Format("15:04:05"))
	fmt.Println("YYYY-MM-DD HH:MM:SS:", now.Format("2006-01-02 15:04:05"))
	//outPut looks like
	//Default format: 2023-10-09 14:05:06.123456789 -0700 PDT m=+0.000000000
	//YYYY-MM-DD: 2023-10-09
	//MM/DD/YYYY: 10/09/2023
	//HH:MM:SS: 14:05:06
	//YYYY-MM-DD HH:MM:SS: 2023-10-09 14:05:06
}

// 2- Parsing Times
// Use time.Parse or time.ParseInLocation to convert a string to a time.Time.
func parseTime() {
	timeStr := "2023-10-09 14:05:06"
	layout := "2006-01-02 15:04:05"
	time, err := time.Parse(layout, timeStr) //this func return the time and the error
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println("Parsed time:", time)
}

// Time Zones
// By default, time.Now() returns the local time. To work with UTC or other time zones:
func timeZone() {
	utcNow := time.Now()
	fmt.Println("UTC time:", utcNow)

	//locations
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}
	nowInNY := time.Now().In(loc)
	fmt.Println("Time in New York:", nowInNY)

}

// /Time Arithmetic
func timeArithmetic() {
	//Adding time
	future := time.Now().Add(2 * time.Hour)
	fmt.Println("Two hours from now:", future)
	//sub
	past := time.Now().Add(-24 * time.Hour)
	fmt.Println("24 hours ago:", past)
	//Calculating Durations Between Times:
	start := time.Now()
	// ... some operations ...
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("Duration:", duration)

	//Comparing Times
	t1 := time.Now()
	t2 := t1.Add(10 * time.Minute)

	fmt.Println("t1 before t2?", t1.Before(t2))  // true
	fmt.Println("t2 after t1?", t2.After(t1))    // true
	fmt.Println("t1 equal to t2?", t1.Equal(t2)) // false

}

// Sleeping and Timing
func sleep() {
	time.Sleep(5 * time.Second)
}
