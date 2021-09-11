package dao

import (
	"fmt"
	"net/http"

	"github.com/petrostrak/yourTrustedMonitoringAndControlPartner/utils"
)

var (
	PeriodicTaskDao periodicTaskDaoService
)

type periodicTaskDaoService interface {
	GetAllPeriodicTasks(string, string, string, string) (*PeriodicTask, *utils.ApplicationError)
}

func init() {
	PeriodicTaskDao = &periodicTaskDao{}
}

type periodicTaskDao struct{}

func (pd *periodicTaskDao) GetAllPeriodicTasks(t1, t2, period, tz string) (*PeriodicTask, *utils.ApplicationError) {
	// Logic to implement if there was a DB
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("%s is greater that %s", t1, t2),
		StatusCode: http.StatusNotFound,
		Code:       "Unsupported period",
	}
}
