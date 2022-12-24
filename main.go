package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var temp Season
var givenYear int
var lastResultsFromGivenYear RaceResults

var raceData Race
var seasonData RaceTable

func getFallBack(dataType string) (data []byte, err error) {
	switch dataType {
	case "calendar":
		data, err = ioutil.ReadFile("./data/calendar_2023.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)

		}
		err = json.Unmarshal(data, &temp)
		seasonData = temp.Data.RaceTable
	case "results":
		data, err = ioutil.ReadFile("./data/round_2022.json")
		if err != nil {
			log.Fatal("Error when opening file: ", err)

		}
		err = json.Unmarshal(data, &lastResultsFromGivenYear)
		raceData = lastResultsFromGivenYear.Data.RaceTable.Races[0]
	}

	return data, err
}

func (calendar *Season) updateCalendar(year int) {

	// var body []byte

	// url := fmt.Sprintf("http://ergast.com/api/f1/%d.json", year)
	// client := &http.Client{}

	// req, _ := http.NewRequest("GET", url, nil)

	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	getFallBack("calendar")
	// }
	// fmt.Println("[CALENDAR] Successfully reached the API")
	// defer res.Body.Close()

	// body, err = ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	getFallBack("calendar")
	// }

	// err = json.Unmarshal(body, &temp)
	// if err != nil {
	// 	fmt.Println(err)
	// 	getFallBack("calendar")
	// }
	getFallBack("calendar")
}

func (race *Race) correctId() {
	for i := 0; i < len(race.Results); i++ {
		race.Results[i].Id = i + 1
	}
}

func main() {
	getCountryData()
	givenYear = 2023
	temp.updateCalendar(givenYear)
	lastResultsFromGivenYear.getResults(givenYear, 1)

	raceData.correctId()
	seasonData.FormatTime()
	seasonData.correctCountry()

	// Handles css files:
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	// Handles routing:
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/resultats", resultsHandler)
	http.HandleFunc("/equipes", teamsHandler)
	http.HandleFunc("/calendrier", calendarHandler)

	// Launches the server:
	preferredPort := ":8080"
	fmt.Printf("Starting server at port %v\n", preferredPort)
	if err := http.ListenAndServe(preferredPort, nil); err != nil {
		log.Fatal(err)
	}
}
