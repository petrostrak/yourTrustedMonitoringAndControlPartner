package utils

import (
	"testing"
)

func TestCheckInvocationPoint(t *testing.T) {
	expectation := true
	actual := CheckInvocationPoint("20060102T150405Z")

	if actual != expectation {
		t.Errorf("Expected %v but got %v", expectation, actual)
	}
}
