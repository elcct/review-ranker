package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	created := "12th July 12:04"
	expected := time.Date(0, 7, 12, 12, 4, 0, 0, time.UTC)

	res, err := parseDate(created)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestStringToSource(t *testing.T) {
	var tests = map[string]Source{
		"solicited":   SourceSolicited,
		"unsolicited": SourceUnsolicited,
		"monkey":      SourceMonkey,
	}

	for test := range tests {
		res, err := stringToSource(test)
		assert.Nil(t, err)
		assert.Equal(t, tests[test], res)
	}
}

func TestStringToWords(t *testing.T) {
	var tests = map[string]int{
		"50 words":  50,
		"1 word":    1,
		"150 words": 150,
	}

	for test := range tests {
		res, err := stringToWords(test)
		assert.Nil(t, err)
		assert.Equal(t, tests[test], res)
	}
}

func TestStringToRating(t *testing.T) {
	var tests = map[string]int{
		"*":     1,
		"**":    2,
		"***":   3,
		"****":  4,
		"*****": 5,
	}

	for test := range tests {
		res, err := stringToRating(test)
		assert.Nil(t, err)
		assert.Equal(t, tests[test], res)
	}
}
