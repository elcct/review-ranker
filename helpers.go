package main

import (
	"strings"
	"time"
)

// parseDate transforms the input string to be suitable for time.Parse method
// and then uses that method to convert date to time.Time type
func parseDate(date string) (time.Time, error) {
	cleanDate := strings.Replace(date, "nd", "", -1)
	cleanDate = strings.Replace(cleanDate, "st", "", -1)
	cleanDate = strings.Replace(cleanDate, "rd", "", -1)
	cleanDate = strings.Replace(cleanDate, "th", "", -1)

	return time.Parse("2 January 15:04", cleanDate)
}
