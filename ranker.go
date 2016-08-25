package main

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
