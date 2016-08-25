package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	// ErrSourceUnknown means string has no known Source representation
	ErrSourceUnknown = errors.New("Source unknown")
	// ErrInvalidWords means string containing number of words is invalid
	ErrInvalidWords = errors.New("Can't parse words value")
)

// parseDate transforms the input string to be suitable for time.Parse method
// and then uses that method to convert date to time.Time type
// TODO: optimize this :)
func parseDate(date string) (time.Time, error) {
	fields := strings.Fields(date)
	cleanDate := strings.Replace(fields[0], "nd", "", 1)
	cleanDate = strings.Replace(cleanDate, "st", "", 1)
	cleanDate = strings.Replace(cleanDate, "rd", "", 1)
	cleanDate = strings.Replace(cleanDate, "th", "", 1)

	return time.Parse("2 January 15:04", fmt.Sprintf("%s %s %s", cleanDate, fields[1], fields[2]))
}

func stringToSource(s string) (Source, error) {
	switch s {
	case "solicited":
		return SourceSolicited, nil
	case "unsolicited":
		return SourceUnsolicited, nil
	case "monkey":
		return SourceMonkey, nil
	}
	return SourceMonkey, ErrSourceUnknown
}

func stringToWords(s string) (int, error) {
	f := strings.Fields(s)
	if len(f) > 1 {
		i, err := strconv.Atoi(f[0])
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, ErrInvalidWords
}

func stringToRating(s string) (int, error) {
	return len(s), nil
}
