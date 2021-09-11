package utils

// ApplicationError is the struct responsible for errors
type ApplicationError struct {
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
	Message    string `json:"desc"`
}
