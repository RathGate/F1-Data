package f1api

type JsonContainer struct {
	JsonContent JsonContent `json:"MRData"`
}

type JsonContent struct {
	Season Season `json:"RaceTable"`
}

type Season struct {
	Year  string `json:"season"`
	Races []Race `json:"Races"`
}

type Race struct {
	Id         string   `json:"round"`
	Name       string   `json:"raceName"`
	Circuit    Circuit  `json:"circuit"`
	Date       string   `json:"date"`
	Time       string   `json:"time"`
	Qualifying Time     `json:"Qualifying"`
	Sprint     Time     `json:"Sprint"`
	Results    []Result `json:"Results"`
}

type Circuit struct {
	Name     string   `json:"circuitName"`
	Location Location `json:"Location"`
}

type Location struct {
	Locality string `json:"locality"`
	Country  string `json:"country"`
}

type Result struct {
	Id          string      `json:"position"`
	Position    string      `json:"positionText"`
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
