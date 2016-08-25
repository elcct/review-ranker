package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"strings"
	"testing"
)

// epsilon tells precision error of float type
// we use it for comparison of float numbers
var epsilon = math.Nextafter(1, 2) - 1

func TestNoopRankerExists(t *testing.T) {
	r := &NoopRanker{}
	assert.NotNil(t, r)
	var i Ranker = r
	assert.NotNil(t, i)
}

func TestLotsToSayRanker(t *testing.T) {
	r, err := NewReviewFromString("13th July 15:04, Jon1, unsolicited, CY8‐IPK, 150 words, ***")
	assert.Nil(t, err)

	ranker := &LotsToSayRanker{}
	err = ranker.Rank(r)
	assert.Nil(t, err)

	p := GetProfessional("Jon1")
	assert.InEpsilon(t, 100.0-100.0*0.005, p.GetScore(), epsilon)
}

func TestBurstRanker(t *testing.T) {
	r1, err := NewReviewFromString("13th July 15:04, Jon2, unsolicited, CY8‐IPK, 150 words, ***")
	assert.Nil(t, err)
	r2, err := NewReviewFromString("13th July 15:04, Jon2, unsolicited, CY8‐IPK, 150 words, ***")
	assert.Nil(t, err)
	r3, err := NewReviewFromString("13th July 15:14, Jon2, unsolicited, CY8‐IPK, 150 words, ***")
	assert.Nil(t, err)

	p := GetProfessional("Jon2")

	ranker := &BurstRanker{}
	err = ranker.Rank(r1)
	assert.Nil(t, err)
	err = ranker.Rank(r2)
	assert.InEpsilon(t, 60, p.GetScore(), epsilon)
	assert.Nil(t, err)
	err = ranker.Rank(r3)
	assert.Nil(t, err)
	assert.InEpsilon(t, 40, p.GetScore(), epsilon)
}

func TestSameDeviceRanker(t *testing.T) {
	r1, err := NewReviewFromString("13th July 15:04, Jon3, unsolicited, CY8‐IPK, 150 words, ***")
	assert.Nil(t, err)
	r2, err := NewReviewFromString("13th July 15:04, Jon3, unsolicited, CY8‐IPK, 150 words, ***")
	assert.Nil(t, err)

	p := GetProfessional("Jon3")

	ranker := &SameDeviceRanker{}
	err = ranker.Rank(r1)
	assert.Nil(t, err)
	err = ranker.Rank(r2)
	assert.Nil(t, err)

	assert.InEpsilon(t, 100.0-100.0*0.3, p.GetScore(), epsilon)
}

func TestAllStarRanker(t *testing.T) {
	r1, err := NewReviewFromString("13th July 15:04, Jon4, unsolicited, CY8‐IPK, 150 words, *****")
	assert.Nil(t, err)
	r2, err := NewReviewFromString("13th July 15:04, Jon4, unsolicited, CY8‐IPK, 150 words, *")
	assert.Nil(t, err)
	r3, err := NewReviewFromString("13th July 15:04, Jon4, unsolicited, CY8‐IPK, 150 words, *")
	assert.Nil(t, err)
	r4, err := NewReviewFromString("13th July 15:04, Jon4, unsolicited, CY8‐IPK, 150 words, *****")
	assert.Nil(t, err)

	p := GetProfessional("Jon4")

	ranker := &AllStarRanker{}
	err = ranker.Rank(r1)
	assert.Nil(t, err)
	assert.InEpsilon(t, 98, p.GetScore(), epsilon)

	err = ranker.Rank(r2)
	assert.Nil(t, err)
	err = ranker.Rank(r3)
	assert.Nil(t, err)
	err = ranker.Rank(r4)
	assert.Nil(t, err)
	assert.InEpsilon(t, 90, p.GetScore(), epsilon)
}

func TestSolicitedRanker(t *testing.T) {
	r, err := NewReviewFromString("13th July 15:04, Jon5, solicited, CY8‐IPK, 150 words, ***")
	assert.Nil(t, err)

	p := GetProfessional("Jon5")
	p.SetScore(50)

	ranker := &SolicitedRanker{}
	err = ranker.Rank(r)
	assert.Nil(t, err)

	assert.InEpsilon(t, 53.0, p.GetScore(), epsilon)
}

func TestReviewRanking(t *testing.T) {
	var rankers = []Ranker{}
	rankers = append(rankers, &LotsToSayRanker{})
	rankers = append(rankers, &BurstRanker{})
	rankers = append(rankers, &SameDeviceRanker{})
	rankers = append(rankers, &AllStarRanker{})
	rankers = append(rankers, &SolicitedRanker{})

	inputs := strings.Split(reviews, "\n")
	outputs := strings.Split(outcomes, "\n")
	for i, input := range inputs {
		r, err := NewReviewFromString(input)
		assert.Nil(t, err)
		assert.NotNil(t, r)

		if r.Source != SourceMonkey {
			for i := range rankers {
				rankers[i].Rank(r)
			}
		}

		assert.Equal(t, outputs[i], r.String())
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
Alert: Jon has been de-activated due to a low trusted review score`
