package main

// AdTypes represents different types of job ads
type AdTypes int

const (
	classic AdTypes = iota
	standout
	premium
)

// Item represents a job ad
type Item struct {
	adType AdTypes
}

// StringToAdTypes converts a string to an adtype
func StringToAdTypes(s string) AdTypes {
	switch s {
	case "classic":
		return classic
	case "standout":
		return standout
	case "premium":
		return premium
	}
	// TODO: what to do if incorrect ad type passed in
	return classic
}
