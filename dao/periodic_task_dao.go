package dao

import (
	"fmt"
	"net/http"
	"petrostrak/yourTrustedMonitoringAndControlPartner/utils"
)

var (
	PeriodicTaskDao periodicTaskDaoService
)

type periodicTaskDaoService interface {
	GetInvocationPoints(string, string) (*PeriodicTask, *utils.ApplicationError)
}

func init() {
	PeriodicTaskDao = &periodicTaskDao{}
}

type periodicTaskDao struct{}

func (pd *periodicTaskDao) GetInvocationPoints(t1, t2 string) (*PeriodicTask, *utils.ApplicationError) {
	// Logic
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("%s is greater that %s", t1, t2),
		StatusCode: http.StatusNotFound,
		Code:       "Unsupported period",
	}
}
