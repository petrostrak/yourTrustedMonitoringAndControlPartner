package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"
)

type respond struct {
	Status int
	Body   interface{}
}

// Respond is the function that returns a successful JSON
func Respond(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	resp := &respond{
		Status: status,
		Body:   body,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatalf("JSON decoding failed: %s", err)
	}

}

// RespondError is the function that returns an ApplicationError
func RespondError(w http.ResponseWriter, err *ApplicationError) {
	w.Header().Set("Content-Type", "application/json")

	resp := &respond{
		Status: err.StatusCode,
		Body:   err,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Fatalf("JSON decoding failed: %s", err)
	}
}

// CheckInvocationPoint checks if invocation points are in the correct format with the use
// of a regular expression
// ^[1-9]\d{3}\d{2}\d{2}T\d{2}\d{2}\d{2}Z$
func CheckInvocationPoint(t string) bool {
	return regexp.MustCompile(`^[1-9]\d{3}\d{2}\d{2}T\d{2}\d{2}\d{2}Z$`).MatchString(t)
}

// CheckInvocationSequence checks if invocation points are in the correct time sequence
func CheckInvocationSequence(t1, t2, layout string) bool {
	ts1, err := time.Parse(layout, t1)
	if err != nil {
		fmt.Println(err)
	}

	ts2, err := time.Parse(layout, t2)
	if err != nil {
		fmt.Println(err)
	}

	return ts1.Before(ts2)
}

// ParseStringToTime receives a string and parses it to time.Time
func ParseStringToTime(layout, invocationPoint string) (*time.Time, *ApplicationError) {
	t1, parseErr := time.Parse(layout, invocationPoint)
	if parseErr != nil {
		err := &ApplicationError{
			Message:    "cannot parse invocation points",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		return nil, err
	}

	return &t1, nil

}
