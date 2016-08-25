package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNoopRankerExists(t *testing.T) {
	r := &NoopRanker{}
	assert.NotNil(t, r)
	var i Ranker = r
	assert.NotNil(t, i)
}

func TestReviewRanking(t *testing.T) {
	inputs := strings.Split(reviews, "\n")
	for _, input := range inputs {
		r, err := NewReviewFromString(input)
		assert.Nil(t, err)
		assert.NotNil(t, r)

	}

}

var reviews = `12th July 12:04, Jon, solicited, LB3‐TYU, 50 words, *****
12th July 12:05, Jon, unsolicited, KB3‐IKU, 20 words, **
13th July 15:04, Jon, unsolicited, CY8‐IPK, 150 words, ***
15th July 10:04, Jon, solicited, BB4‐IPK, 40 words, *****
15th July 15:09, Jon, monkey
29th August 10:04, Jon, solicited, LX2‐IPK, 70 words, ****
2nd September 10:04, Jon, solicited, KB3‐IKU, 50 words, ****
2nd September 10:04, Jon, solicited, AN9‐IPK, 90 words, **`

var outcomes = `Info: Jon has a trusted review score of 100
Info: Jon has a trusted review score of 80
Info: Jon has a trusted review score of 79.5
Info: Jon has a trusted review score of 74.5
Could not read review summary data
Info: Jon has a trusted review score of 77.5
Warning: Jon has a trusted review score of 50.5
Alert: Jon has been de‐activated due to a low trusted review score`
