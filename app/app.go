package app

import (
	"fmt"
	"net/http"
	"os"
)

var (
	mux *http.ServeMux
)

func init() {
	mux = http.NewServeMux()
}

func parsePort() string {
	// go run *.go -p 4000
	// p := flag.Int("p", 8080, "Give a port to listen")
	// flag.Parse()
	// port := strconv.Itoa(*p)

	if len(os.Args) != 2 {
		fmt.Println("give a port to listen to")
		os.Exit(0)
	}
	port := os.Args[1]

	fmt.Printf("Listening on port %s\n", port)
	return ":" + port
}

func StartApp() {
	mapURLs()

	if err := http.ListenAndServe(parsePort(), mux); err != nil {
		panic(err)
	}
}
