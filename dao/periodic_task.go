package dao

import "time"

// PeriodicTask struct is the collection of attributes for periodic task
type PeriodicTask struct {
	Period   time.Duration `json:"period"`
	Timezone string        `json:"tz"`
	InvocationPoints
	Timestamps []string `json:"timestamps"`
}

type InvocationPoints struct {
	T1 string `json:"t1"`
	T2 string `json:"t2"`
}
