package f1api

import (
	"encoding/json"
	"os"
)

var Countries []Country

type jsonCountry struct {
	Countries []Country `json:"flags"`
}
type Country struct {
	Code        string
	Nationality string
	Country     string
}

func SetCountryData() (err error) {
	var temp jsonCountry
	data, err := os.ReadFile("./data/flags.json")

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	Countries = temp.Countries
	return nil
}

func (race Race) RetrieveFlag() string {
	for _, element := range Countries {
		if race.Circuit.Location.Country == element.Country || race.Circuit.Location.Country == element.Nationality {
			return element.Code
		}
	}
	return "unknown"
}

func (season *Season) CorrectCountry() {
	for i := 0; i < len(season.Races); i++ {

		switch season.Races[i].Circuit.Location.Country {
		case "UAE":
			season.Races[i].Circuit.Location.Country = "United Arab Emirates"

		case "USA":
			season.Races[i].Circuit.Location.Country = "United States of America"

		case "United States":
			season.Races[i].Circuit.Location.Country = "United States of America"

		case "UK":
			season.Races[i].Circuit.Location.Country = "Great Britain"
		}
	}
}
