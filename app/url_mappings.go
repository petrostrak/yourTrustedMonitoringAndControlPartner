package app

import "petrostrak/yourTrustedMonitoringAndControlPartner/controllers"

func mapURLs() {
	mux.HandleFunc("/ptlist", controllers.GetPeriodicTask)
}
