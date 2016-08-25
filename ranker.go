package main

import (
	"math"
	"time"
)

// Ranker is an interface for Rankers
type Ranker interface {
	Rank(*Review) error
}

// NoopRanker does nothing
type NoopRanker struct {
}

// Rank does nothing
func (r *NoopRanker) Rank(review *Review) (err error) {
	return
}

// LotsToSayRanker knocks off 0.5% of score when review has >100 words
type LotsToSayRanker struct {
}

// Rank performs ranking operation
func (r *LotsToSayRanker) Rank(review *Review) (err error) {
	if review.Words > 100 {
		review.Professional.AdjustScore(-0.005)
	}
	return
}

// BurstRanker knocks off 40% of score if review comes in the same frame
type BurstRanker struct {
	last time.Time
}

// Rank performs ranking operation
func (r *BurstRanker) Rank(review *Review) (err error) {
	if review.CreatedAt == r.last {
		review.Professional.AdjustScore(-0.4)
	} else {
		diff := review.CreatedAt.Sub(r.last)
		if math.Abs(diff.Hours()) <= 1 {
			review.Professional.AdjustScore(-0.2)
		}
	}

	r.last = review.CreatedAt
	return
}

// SameDeviceRanker knocks off 30% of score if review comes from the same device
type SameDeviceRanker struct {
	devices map[string]bool
}

// Rank performs ranking operation
func (r *SameDeviceRanker) Rank(review *Review) (err error) {
	if _, exists := r.devices[review.Device]; exists {
		review.Professional.AdjustScore(-0.3)
	} else {
		if r.devices == nil {
			r.devices = map[string]bool{}
		}
		r.devices[review.Device] = true
	}
	return
}

// AllStarRanker knocks off 2% of score if review comes with 5 star rating
// it quadruples the penalty if rating average is below 3.5
type AllStarRanker struct {
	average float32
	samples int32
}

// Rank performs ranking operation
func (r *AllStarRanker) Rank(review *Review) (err error) {
	if review.Rating == 5 {
		if r.average != 0.0 && r.average < 3.5 {
			review.Professional.AdjustScore(-0.08)
		} else {
			review.Professional.AdjustScore(-0.02)
		}
	}

	if r.average == 0.0 {
		r.average = float32(review.Rating)
		r.samples = 1
	} else {
		// Calculate moving average
		r.samples++
		r.average -= r.average / float32(r.samples)
		r.average += float32(review.Rating) / float32(r.samples)
	}
	return
}

// SolicitedRanker adds 3% of score when review is solicited
type SolicitedRanker struct {
}

// Rank performs ranking operation
func (r *SolicitedRanker) Rank(review *Review) (err error) {
	if review.Source == SourceSolicited {
		review.Professional.AdjustScore(0.03)
	}
	return
}
