package main

import (
	"html/template"
	"net/http"
	"projet-web/packages/f1api"
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

	if r.Method == "POST" {
		w.Write([]byte("received"))
		return
	}
	files := []string{"templates/results.html", "templates/_results-table.html"}
	template := template.Must(template.ParseFiles(files...))
	template.Execute(w, latestData)
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/teams.html"))
	template.Execute(w, nil)
}
