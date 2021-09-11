package services

import (
	"github.com/petrostrak/yourTrustedMonitoringAndControlPartner/dao"
	"github.com/petrostrak/yourTrustedMonitoringAndControlPartner/utils"
)

type periodicTaskService struct{}

var (
	PeriodicTaskService periodicTaskService
)

// GetInvocationPoints receives the invocation points as parameters from the controller
func (ps *periodicTaskService) GetAllPeriodicTasks(t1, t2, period, tz string) (*dao.PeriodicTask, *utils.ApplicationError) {
	// then in calls DAO to check in the DB
	// and returns the results to controller
	return dao.PeriodicTaskDao.GetAllPeriodicTasks(t1, t2, period, tz)
}
