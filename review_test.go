package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReviewExist(t *testing.T) {
	r := &Review{}
	assert.NotNil(t, r)
}

func TestCreateAndPrintReview(t *testing.T) {
	created, err := parseDate("12th July 12:04")
	assert.Nil(t, err)

	r := NewReview(
		created,
		GetProfessional("Jon100"),
		SourceSolicited,
		"LB3-TYU",
		50,
		5,
	)
	assert.NotNil(t, r)

	// Before processed by ranker, review will have 100 score
	s := "Info: Jon100 has a trusted review score of 100"
	assert.Equal(t, s, r.String())
}

func TestReviewFromString(t *testing.T) {
	source := "12th July 12:04, Jon101, solicited, LB3‐TYU, 50 words, *****"
	r, err := NewReviewFromString(source)
	assert.Nil(t, err)
	assert.NotNil(t, r)

	// Before processed by ranker, review will have 100 score
	s := "Info: Jon101 has a trusted review score of 100"
	assert.Equal(t, s, r.String())
}
