package main

import (
	"fmt"
	"log"
	"net/http"
	"projet-web/packages/f1api"
)

const CURRENT_YEAR = 2023

var latestData Data
var currentCalendar f1api.Season
var allRaces []f1api.Race

type Data struct {
	SeasonData f1api.Season
	RaceData   f1api.Race
}

func main() {

	initData()
	// Handles css files:
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles routing:
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/results", resultsHandler)
	http.HandleFunc("/teams", teamsHandler)
	http.HandleFunc("/calendar", calendarHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}

}
