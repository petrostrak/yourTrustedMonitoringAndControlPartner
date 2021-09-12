package app

import "github.com/petrostrak/yourTrustedMonitoringAndControlPartner/controllers"

func mapURLs() {
	mux.HandleFunc("/ptlist", controllers.GetAllPeriodicTasks)
}
