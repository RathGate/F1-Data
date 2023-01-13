package main

import (
	"html/template"
	"net/http"
	"projet-web/packages/f1api"
	"regexp"
	"strconv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/index.html"))
	template.Execute(w, nil)
}

func calendarHandler(w http.ResponseWriter, r *http.Request) {
	test, err := template.New("calendar.html").Funcs(template.FuncMap{
		"formatTime": func(time string) string {
			return time[:5]
		},
		"formatDate": f1api.FormatTime,
	}).ParseFiles("templates/calendar.html")
	if err != nil {
		panic(err)
	}

	err = test.Execute(w, currentCalendar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" && r.FormValue("requestedYear") != "" {
		str_id := regexp.MustCompile(`\d+`).FindString(r.FormValue("requestedYear"))
		id, _ := strconv.Atoi(str_id)

		var tempData = Data{
			SeasonData: latestData.SeasonData,
			RaceData:   allRaces[id-1],
		}

		test, err := template.New("results-table").Funcs(template.FuncMap{
			"retrieveLast": func(arr []string) string {
				if len(arr) == 0 {
					return ""
				}
				return arr[len(arr)-1]
			},
			"formatTime": func(time string) string {
				return time[:5]
			},
			"formatDate": f1api.FormatTime,
		}).ParseFiles("templates/_results-table.html")
		if err != nil {
			panic(err)
		}

		err = test.Execute(w, tempData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	files := []string{"templates/_results-table.html", "templates/results.html"}
	test, err := template.New("results.html").Funcs(template.FuncMap{
		"retrieveLast": func(arr []string) string {
			if len(arr) == 0 {
				return ""
			}
			return arr[len(arr)-1]
		},
		"formatTime": func(time string) string {
			return time[:5]
		},
		"formatDate": f1api.FormatTime,
	}).ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	err = test.Execute(w, latestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/teams.html"))
	template.Execute(w, nil)
}
