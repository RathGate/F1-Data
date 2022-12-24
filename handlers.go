package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	template := template.Must(template.ParseFiles("templates/index.html"))
	template.Execute(w, nil)
}
func calendarHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/calendar.html"))
	template.Execute(w, seasonData)
}
func resultsHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/results.html"))
	template.Execute(w, raceData.Results)
}
func teamsHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/teams.html"))
	template.Execute(w, nil)
}

type Season struct {
	Data Data `json:"MRData"`
}

type Data struct {
	RaceTable RaceTable `json:"RaceTable"`
}

type RaceTable struct {
	Year  string `json:"season"`
	Races []Race `json:"Races"`
}
type Race struct {
	Id      string   `json:"round"`
	Name    string   `json:"raceName"`
	Circuit Circuit  `json:"circuit"`
	Date    string   `json:"date"`
	Time    string   `json:"time"`
	FP1     Time     `json:"FirstPractice"`
	Results []Result `json:"Results"`
}

type Circuit struct {
	Name     string   `json:"circuitName"`
	Location Location `json:"Location"`
}

type Location struct {
	Locality string `json:"locality"`
	Country  string `json:"country"`
}

var months = []string{"", "JAN", "FEV", "MARS", "AVR", "MAI", "JUIN", "JUIL", "AOUT", "SEPT", "OCT", "NOV", "DEC"}

func (race Race) GetTime() map[string][]string {
	interval := make(map[string][]string)
	startingDate := DateToArr(race.FP1.Date)
	endingDate := DateToArr(race.Date)

	interval["months"] = append(interval["months"], months[startingDate[1]])
	if startingDate[1] != endingDate[1] {
		interval["months"] = append(interval["months"], months[endingDate[1]])
	}
	interval["days"] = append(interval["days"], fmt.Sprintf("%02d", startingDate[2]), fmt.Sprintf("%02d", endingDate[2]))

	return interval
}

func DateToArr(date string) (arr []int) {
	temp := strings.Split(date, "-")
	for _, element := range temp {
		nb, _ := strconv.Atoi(element)
		arr = append(arr, nb)
	}
	return arr
}
