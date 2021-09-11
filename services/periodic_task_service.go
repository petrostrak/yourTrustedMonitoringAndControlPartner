package services

import (
	"github.com/petrostrak/yourTrustedMonitoringAndControlPartner/dao"
	"github.com/petrostrak/yourTrustedMonitoringAndControlPartner/utils"
)

type periodicTaskService struct{}

var (
	PeriodicTaskService periodicTaskService
)

func (ps *periodicTaskService) GetInvocationPoints(t1, t2 string) (*dao.PeriodicTask, *utils.ApplicationError) {
	return dao.PeriodicTaskDao.GetInvocationPoints(t1, t2)
}
