package main

// JobAd represents different types of job ads
type JobAd int

const (
	classic JobAd = iota
	standout
	premium
)

// Item represents a job ad
type Item struct {
	id JobAd
}
