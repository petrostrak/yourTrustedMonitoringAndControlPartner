package app

func mapURLs() {
	mux.HandleFunc("/ptlist", controllers.GetPeriodicTask)
}
