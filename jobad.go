package main

// AdType represents different types of job ads
type AdType int

const (
	classic AdType = iota
	standout
	premium
)

// Item represents a job ad
type Item struct {
	adType AdType
}
