package main

// Truncate a float to two levels of precision
func Truncate(some float64) float64 {
	return float64(int(some*100)) / 100
}
