package services

import (
	"petrostrak/yourTrustedMonitoringAndControlPartner/dao"
	"petrostrak/yourTrustedMonitoringAndControlPartner/utils"
)

type periodicTaskService struct{}

var (
	PeriodicTaskService periodicTaskService
)

func (ps *periodicTaskService) GetInvocationPoints(t1, t2 string) (*dao.PeriodicTask, *utils.ApplicationError) {
	return dao.PeriodicTaskDao.GetInvocationPoints(t1, t2)
}
