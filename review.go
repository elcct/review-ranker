package main

import (
	"fmt"
	"strings"
	"time"
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

const (
	// StatusLevelInfo defines minimum score for Info level
	StatusLevelInfo float32 = 70
	// StatusLevelWarning defines minimum score for Warning level
	StatusLevelWarning float32 = 50
)

// Status is a type defining status of the reviwe
type Status int

const (
	// StatusInfo means review is valid
	StatusInfo Status = iota
	// StatusWarning means review should be investigated
	StatusWarning
	// StatusAlert means review should not be trusted
	StatusAlert
	// StatusInvalid means review is not valid
	StatusInvalid
)

func (s Status) String() string {
	switch s {
	case StatusInfo:
		return "Info"
	case StatusWarning:
		return "Warning"
	case StatusAlert:
		return "Alert"
	case StatusInvalid:
		return "Invalid"
	}
	return "Unknown"
}

// Review defines review attributes
type Review struct {
	CreatedAt time.Time `json:"time"`
	Author    string    `json:"author"`
	Source    Source    `json:"source"`
	Device    string    `json:"device"`
	Words     int       `json:"words"`
	Rating    int       `json:"rating"`

	status Status  `json:"status"`
	score  float32 `json:"score"`
}

// NewReview creates new instance of Review
func NewReview(createdAt time.Time,
	author string,
	source Source,
	device string,
	words int,
	rating int) *Review {
	r := &Review{
		CreatedAt: createdAt,
		Author:    author,
		Source:    source,
		Device:    device,
		Words:     words,
		Rating:    rating,
	}

	r.SetScore(100)
	return r
}

// NewReviewFromString creates new review from string
// string should be in following format:
// 12th July 12:04, Jon, solicited, LB3‐TYU, 50 words, *****
func NewReviewFromString(input string) (*Review, error) {
	// split string by ,
	fields := strings.FieldsFunc(input, func(c rune) bool {
		return c == ','
	})

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

	words, err := stringToWords(fields[4])
	if err != nil {
		return nil, err
	}

	rating, err := stringToRating(fields[5])
	if err != nil {
		return nil, err
	}

	r := NewReview(
		createdAt,
		fields[1],
		source,
		fields[3],
		words,
		rating,
	)

	return r, nil
}

// SetScore sets score of the review and updates review status accordingly
func (r *Review) SetScore(score float32) {
	r.score = score
	if r.score >= StatusLevelInfo {
		r.status = StatusInfo
		return
	}
	if r.score >= StatusLevelWarning {
		r.status = StatusWarning
		return
	}
	r.status = StatusAlert
}

// GetScore return current review score
func (r Review) GetScore() float32 {
	return r.score
}

// GetStatus returns current review status
func (r Review) GetStatus() Status {
	return r.status
}

func (r Review) String() string {
	if r.Source == SourceMonkey {
		return "Could not read summary data"
	}
	switch r.status {
	case StatusInfo:
		fallthrough
	case StatusWarning:
		return fmt.Sprintf("%s: %s has a trusted review score of %2.f", r.status, r.Author, r.score)
	case StatusAlert:
		return fmt.Sprintf("%s: %s has been de-activated due to a low trusted review score", r.status, r.Author)
	}

	return ""
}
