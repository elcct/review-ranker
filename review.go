package main

import (
	"errors"
	"strings"
	"time"
)

var (
	// ErrInvalidString means that we can't create Review from given input string
	ErrInvalidString = errors.New("Review couldn't be parsed, invalid input string.")
)

// Source is a type that defines where review comes from
type Source int

const (
	// SourceSolicited means review is coming from the person invited by a professional
	SourceSolicited Source = iota
	// SourceUnsolicited means review was from not invited person
	SourceUnsolicited
	// SourceMonkey means something is wrong and this review should be ignored
	SourceMonkey
)

// Review defines review attributes
type Review struct {
	CreatedAt    time.Time     `json:"time"`
	Professional *Professional `json:"professional"`
	Source       Source        `json:"source"`
	Device       string        `json:"device"`
	Words        int           `json:"words"`
	Rating       int           `json:"rating"`
}

// NewReview creates new instance of Review
func NewReview(createdAt time.Time,
	professional *Professional,
	source Source,
	device string,
	words int,
	rating int) *Review {
	r := &Review{
		CreatedAt:    createdAt,
		Professional: professional,
		Source:       source,
		Device:       device,
		Words:        words,
		Rating:       rating,
	}
	return r
}

// NewReviewFromString creates new review from string
// string should be in following format:
// 12th July 12:04, Jon, solicited, LB3‐TYU, 50 words, *****
func NewReviewFromString(input string) (r *Review, err error) {
	// split string by ,
	fields := strings.FieldsFunc(input, func(c rune) bool {
		return c == ','
	})

	if len(fields) < 2 {
		return nil, ErrInvalidString
	}

	// clean each field
	for i := range fields {
		fields[i] = strings.TrimSpace(fields[i])
	}

	createdAt, err := parseDate(fields[0])
	if err != nil {
		return nil, err
	}

	source, err := stringToSource(fields[2])
	if err != nil {
		return nil, err
	}

	words := 0
	rating := 0
	device := ""

	professional := GetProfessional(fields[1])

	if source != SourceMonkey {
		words, err = stringToWords(fields[4])
		if err != nil {
			return nil, err
		}

		rating, err = stringToRating(fields[5])
		if err != nil {
			return nil, err
		}

		device = fields[3]
	}

	r = NewReview(
		createdAt,
		professional,
		source,
		device,
		words,
		rating,
	)

	return
}

func (r Review) String() string {
	if r.Source == SourceMonkey {
		return "Could not read review summary data"
	}
	return r.Professional.String()
}
