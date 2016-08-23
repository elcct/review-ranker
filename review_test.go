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
		"Jon",
		Solicited,
		"LB3-TYU",
		50,
		5,
	)
	assert.NotNil(t, r)
}
