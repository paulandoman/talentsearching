package main

import (
	"testing"
)

func TestTruncate(t *testing.T) {
	val := 197.9001
	expected := 197.90
	truncVal := Truncate(val)
	if truncVal != expected {
		t.Errorf("Didn't round down value succesfully, got: %v, expected: %v", truncVal, expected)
	}

	val = 23.255
	expected = 23.25
	truncVal = Truncate(val)
	if truncVal != expected {
		t.Errorf("Didn't round down value succesfully, got: %v, expected: %v", truncVal, expected)
	}

	val = 10.0099
	expected = 10.00
	truncVal = Truncate(val)
	if truncVal != expected {
		t.Errorf("Didn't round down value succesfully, got: %v, expected: %v", truncVal, expected)
	}

}
