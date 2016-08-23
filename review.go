package main

import (
	"time"
)

// Source defines where review comes from
type Source int

const (
	// Solicited means review is coming from the person invited by a professional
	Solicited Source = iota
	// Unsolicited means review was from not invited person
	Unsolicited
	// Monkey means something is wrong and this review should be ignored
	Monkey
)

// Review defines review attributes
type Review struct {
	CreatedAt time.Time `json:"time"`
	Author    string    `json:"author"`
	Source    Source    `json:"source"`
	Device    string    `json:"device"`
	Words     int       `json:"words"`
	Rating    int       `json:"rating"`
}

// NewReview created new instance of Review
func NewReview(createdAt time.Time,
	author string,
	source Source,
	device string,
	words int,
	rating int) *Review {
	return &Review{
		CreatedAt: createdAt,
		Author:    author,
		Source:    source,
		Device:    device,
		Words:     words,
		Rating:    rating,
	}
}
