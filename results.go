package main

import (
	"fmt"
)

type RaceResults struct {
	Data Data `json:"MRData"`
}

type Result struct {
	Id          int
	Points      string      `json:"points"`
	Driver      Driver      `json:"Driver"`
	Laps        string      `json:"laps"`
	Constructor Constructor `json:"Constructor"`
	Time        Time        `json:"Time"`
	Status      string      `json:"status"`
}

type Driver struct {
	Number      string `json:"permanentNumber"`
	Code        string `json:"code"`
	FirstName   string `json:"givenName"`
	FamilyName  string `json:"familyName"`
	Nationality string `json:"nationality"`
}

type Constructor struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
}

type Time struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

func (results *RaceResults) getResults(year, round int) {
	getFallBack("results")
	// _, _ = year, round
	// url := "http://ergast.com/api/f1/2022/21/results.json"

	// client := &http.Client{}
	// req, err := http.NewRequest("GET", url, nil)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer res.Body.Close()

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// err = json.Unmarshal(body, &lastResultsFromGivenYear)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// nique := &lastResultsFromGivenYear.Data.RaceTable.Races[0]
}

func (race *Race) FormatResults() {
	for index, element := range race.Results {
		p := &element.Id
		*p = index
		fmt.Println(index+1, element.Id)
	}
}

func (table *RaceTable) FormatTime() {
	for i := 0; i < len(table.Races); i++ {
		table.Races[i].Time = table.Races[i].Time[:5]
	}
}
