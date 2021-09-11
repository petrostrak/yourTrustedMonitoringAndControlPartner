package controllers

import (
	"fmt"
	"net/http"
	"petrostrak/yourTrustedMonitoringAndControlPartner/dao"
	"petrostrak/yourTrustedMonitoringAndControlPartner/utils"

	"time"
)

const (
	UTC_FORM = "20060102T150405Z"
)

func GetPeriodicTask(w http.ResponseWriter, r *http.Request) {
	// read query parameteres
	queryParams := r.URL.Query()

	period := queryParams.Get("period")
	timezone := queryParams.Get("tz")
	t1, t2 := queryParams.Get("t1"), queryParams.Get("t2")

	// 1h, 1d, 1mo, 1y
	p, err := parsePeriod(period)
	if p == "" || err != nil {
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
// helper function to check period
// Valid periods should be 1h, 1d, 1mo, 1y
func parsePeriod(p string) (string, *utils.ApplicationError) {
	switch p {
	case "1h":
		return p, nil
	case "1d":
		return p, nil
	case "1mo":
		return p, nil
	case "1y":
		return p, nil
	default:
		return "", &utils.ApplicationError{
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

// helper function to check invocation points
func parseInvocationPoints(t1, t2, period string) ([]string, *utils.ApplicationError) {

	timestamps := []string{}

	if utils.CheckInvocationPoint(t1) && utils.CheckInvocationPoint(t2) {

		if utils.CheckInvocationSequence(t1, t2, UTC_FORM) {
			// Valid periods should be 1h, 1d, 1mo, 1y
			switch period {
			case "1h":
				timestamps, err := calculateTimestampsPerHour(t1, t2)
				if err != nil {
					return nil, err
				}

				return timestamps, nil
			case "1d":
				timestamps, err := calculateTimestampsPerDay(t1, t2)
				if err != nil {
					return nil, err
				}

				return timestamps, nil
			case "1mo":
				timestamps, err := calculateTimestampsPerMonth(t1, t2)
				if err != nil {
					return nil, err
				}

				return timestamps, nil
			case "1y":
				timestamps, err := calculateTimestampsPerYear(t1, t2)
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

// calculateTimestampsPerHour appends the timestamps into the slice  for every hour
func calculateTimestampsPerHour(t1, t2 string) ([]string, *utils.ApplicationError) {
	timestamps := []string{}

	ip1, err := utils.ParseStringToTime(UTC_FORM, t1)
	if err != nil {
		return nil, err
	}

	ip2, err := utils.ParseStringToTime(UTC_FORM, t2)
	if err != nil {
		return nil, err
	}

	difference := ip2.Sub(*ip1).Round(60 * time.Minute).Hours()
	timespace := int(difference)

	h := ip1.Hour()
	year, month, day := ip1.Date()

	for i := 0; i < timespace; i++ {
		timestamp := time.Date(year, month, day, h+i, 0, 0, 0, time.UTC)
		timestamps = append(timestamps, timestamp.Round(60*time.Minute).Format(UTC_FORM))
	}

	return timestamps, nil
}

// calculateTimestampsPerDay appends the timestamps into the slice  for every day
func calculateTimestampsPerDay(t1, t2 string) ([]string, *utils.ApplicationError) {
	timestamps := []string{}

	ip1, err := utils.ParseStringToTime(UTC_FORM, t1)
	if err != nil {
		return nil, err
	}

	ip2, err := utils.ParseStringToTime(UTC_FORM, t2)
	if err != nil {
		return nil, err
	}

	difference := ip2.Sub(*ip1).Round(60*time.Minute).Hours() / 24
	timespace := int(difference)

	h := ip1.Hour()
	year, month, day := ip1.Date()

	for i := 0; i < timespace; i++ {
		timestamp := time.Date(year, month, day, h, 0, 0, 0, time.UTC)
		timestamps = append(timestamps, timestamp.AddDate(0, 0, i).Format(UTC_FORM))
	}

	return timestamps, nil
}

// calculateTimestampsPerMonth appends the timestamps into the slice  for every month
func calculateTimestampsPerMonth(t1, t2 string) ([]string, *utils.ApplicationError) {
	timestamps := []string{}

	ip1, err := utils.ParseStringToTime(UTC_FORM, t1)
	if err != nil {
		return nil, err
	}

	ip2, err := utils.ParseStringToTime(UTC_FORM, t2)
	if err != nil {
		return nil, err
	}

	difference := ip2.Sub(*ip1).Round(60*time.Minute).Hours() / 720
	timespace := int(difference)

	h := ip1.Hour()
	year, month, day := ip1.Date()

	for i := 0; i < timespace; i++ {
		timestamp := time.Date(year, month, day, h, 0, 0, 0, time.UTC)
		timestamps = append(timestamps, timestamp.AddDate(0, i, 0).Format(UTC_FORM))
	}

	return timestamps, nil
}

// calculateTimestampsPerYear appends the timestamps into the slice  for every year
func calculateTimestampsPerYear(t1, t2 string) ([]string, *utils.ApplicationError) {
	timestamps := []string{}

	ip1, err := utils.ParseStringToTime(UTC_FORM, t1)
	if err != nil {
		return nil, err
	}

	ip2, err := utils.ParseStringToTime(UTC_FORM, t2)
	if err != nil {
		return nil, err
	}

	difference := ip2.Sub(*ip1).Round(60*time.Minute).Hours() / 8640
	timespace := int(difference)

	h := ip1.Hour()
	year, month, day := ip1.Date()

	for i := 0; i < timespace; i++ {
		timestamp := time.Date(year, month, day, h, 0, 0, 0, time.UTC)
		timestamps = append(timestamps, timestamp.AddDate(0, i, 0).Format(UTC_FORM))
	}

	return timestamps, nil
}
