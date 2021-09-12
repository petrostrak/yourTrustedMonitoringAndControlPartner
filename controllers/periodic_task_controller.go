package controllers

import (
	"fmt"
	"net/http"

	"time"

	"github.com/petrostrak/yourTrustedMonitoringAndControlPartner/dao"
	"github.com/petrostrak/yourTrustedMonitoringAndControlPartner/utils"
)

const (
	UTC_FORM = "20060102T150405Z"

	// periods
	hour  = 60 * time.Minute
	day   = 24 * hour
	week  = 7 * day
	month = 4 * week
	year  = 12 * month
)

func GetAllPeriodicTasks(w http.ResponseWriter, r *http.Request) {
	// read query parameteres
	queryParams := r.URL.Query()

	period := queryParams.Get("period")
	timezone := queryParams.Get("tz")
	t1, t2 := queryParams.Get("t1"), queryParams.Get("t2")

	// 1h, 1d, 1mo, 1y
	p, err := parsePeriod(period)
	if p == -1 || err != nil {
		utils.RespondError(w, err)
		return
	}

	// Europe/Athens
	tz, err := parseTimezone(timezone)
	if tz == "" || err != nil {
		utils.RespondError(w, err)
		return
	}

	invocationPoints, err := parseInvocationPoints(t1, t2, p)
	if err != nil {
		utils.RespondError(w, err)
		return
	}

	// We would call GetAllPeriodicTasks from services to implement MVC if we were
	// working with data persistence.
	//
	// pd, err := services.PeriodicTaskService.GetAllPeriodicTasks(t1, t2, p, tz)
	// if err != nil {
	// 	utils.RespondError(w, err)
	// }

	pd := &dao.PeriodicTask{
		Period:   p,
		Timezone: tz,
		InvocationPoints: dao.InvocationPoints{
			T1: t1,
			T2: t2,
		},
		Timestamps: invocationPoints,
	}

	utils.Respond(w, http.StatusOK, pd)
}

// /ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z
// helper function to parse period and translate it into time.Duration
// Valid periods should be 1h, 1d, 1mo, 1y
func parsePeriod(p string) (time.Duration, *utils.ApplicationError) {
	switch p {
	case "1h":
		return hour, nil
	case "1d":
		return day, nil
	case "1mo":
		return month, nil
	case "1y":
		return year, nil
	default:
		return -1, &utils.ApplicationError{
			Message:    fmt.Sprintf("could not parse period : %s", p),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
	}
}

// helper function to check timezone
func parseTimezone(tz string) (string, *utils.ApplicationError) {
	timeZone, errTZ := time.LoadLocation(tz)
	if errTZ != nil {
		err := &utils.ApplicationError{
			Message:    fmt.Sprintf("could not parse timezone : %s", timeZone),
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}

		return "", err
	}

	return timeZone.String(), nil
}

// parseInvocationPoints checks invocation points and calculates the timestamps, if any
func parseInvocationPoints(t1, t2 string, period time.Duration) ([]string, *utils.ApplicationError) {

	timestamps := []string{}

	if utils.CheckInvocationPoint(t1) && utils.CheckInvocationPoint(t2) {

		if utils.CheckInvocationSequence(t1, t2, UTC_FORM) {
			// Valid periods should be 1h, 1d, 1mo, 1y
			switch period {
			case hour:
				timestamps, err := calculateTimestamps(t1, t2, period)
				if err != nil {
					return nil, err
				}

				return timestamps, nil
			case day:
				timestamps, err := calculateTimestamps(t1, t2, period)
				if err != nil {
					return nil, err
				}

				return timestamps, nil
			case month:
				timestamps, err := calculateTimestamps(t1, t2, period)
				if err != nil {
					return nil, err
				}

				return timestamps, nil
			case year:
				timestamps, err := calculateTimestamps(t1, t2, period)
				if err != nil {
					return nil, err
				}

				return timestamps, nil
			}
		}

		return timestamps, nil
	} else {
		return nil, &utils.ApplicationError{
			Message:    "invocation points do not follow the correct format",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
	}
}

// calculateTimestampsPerHour appends the timestamps into the slice
func calculateTimestamps(t1, t2 string, period time.Duration) ([]string, *utils.ApplicationError) {
	timestamps := []string{}

	ip1, err := utils.ParseStringToTime(UTC_FORM, t1)
	if err != nil {
		return nil, err
	}

	ip2, err := utils.ParseStringToTime(UTC_FORM, t2)
	if err != nil {
		return nil, err
	}

	// The idea behind this algorithm is to take those two invocation points,
	// find the time difference between them, and divide it with the given
	// period in order to find the timestamps
	difference := ip2.Sub(*ip1).Round(60*time.Minute) / period
	timespace := int(difference)

	h := ip1.Hour()
	year, month, day := ip1.Date()

	for i := 0; i < timespace; i++ {
		timestamp := time.Date(year, month, day, h+i, 0, 0, 0, time.UTC)
		timestamps = append(timestamps, timestamp.Round(60*time.Minute).Format(UTC_FORM))
	}

	return timestamps, nil
}
