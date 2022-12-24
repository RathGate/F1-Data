package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var Countries []Country

type CountryData struct {
	Countries []Country `json:"flags"`
}
type Country struct {
	Code        string
	Nationality string
	Country     string
}

func getCountryData() {
	var temp CountryData
	data, err := ioutil.ReadFile("./data/flags.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)

	}
	_ = json.Unmarshal(data, &temp)
	Countries = temp.Countries
}

func (race Race) RetrieveFlag() string {
	for _, element := range Countries {
		if race.Circuit.Location.Country == element.Country || race.Circuit.Location.Country == element.Nationality {
			return element.Code
		}
	}
	return "unknown"
}

func (table *RaceTable) correctCountry() {
	for i := 0; i < len(table.Races); i++ {

		switch table.Races[i].Circuit.Location.Country {
		case "UAE":
			fmt.Println("Hello")
			table.Races[i].Circuit.Location.Country = "United Arab Emirates"

		case "USA":
			table.Races[i].Circuit.Location.Country = "United States of America"

		case "United States":
			table.Races[i].Circuit.Location.Country = "United States of America"

		case "UK":
			table.Races[i].Circuit.Location.Country = "Great Britain"
		}
	}
}
