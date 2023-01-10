package main

import (
	"fmt"
	"log"
	"projet-web/packages/f1api"
)

func initData() {
	err := f1api.SetCountryData()
	if err != nil {
		log.Fatal(err)
	}

	// if err = retrieveInitialAPIData(); err != nil {
	// 	fmt.Println("ERROR: ", err)
	retrieveFallbackData()
	// 	return
	// }
	currentCalendar.CorrectCountry()
	data_initialized = true
	fmt.Println("DATA INITIALIZED SUCCESSFULLY")
}

// func retrieveInitialAPIData() (err error) {
// 	fmt.Println("INITIALIZING DATA...")
// 	// Current season:
// 	temp, err := f1api.APIresponse("https://ergast.com/api/f1/current.json")
// 	if err != nil {
// 		return err
// 	}
// 	latestData.SeasonData = temp.JsonContent.Season

// 	// Latest race from current season:
// 	temp, err = f1api.APIresponse("https://ergast.com/api/f1/current/last/results.json")
// 	if err != nil {
// 		return err
// 	}
// 	latestData.RaceData = temp.JsonContent.Season.Races[0]

// 	// Sets current calendar to 2023 if not already the case:
// 	if year, _ := strconv.Atoi(latestData.SeasonData.Year); year == CURRENT_YEAR {
// 		currentCalendar = latestData.SeasonData
// 	} else {
// 		temp, err := f1api.APIresponse("https://ergast.com/api/f1/2023.json")
// 		if err != nil {
// 			return err
// 		}
// 		currentCalendar = temp.JsonContent.Season
// 	}
// 	return nil
// }

func retrieveFallbackData() {
	fmt.Print("----------\nRetrieving data from local storage..\n\n")
	// 2022 season data:
	temp, err := f1api.GetFallBack("calendar_2022")
	if err != nil {
		log.Fatal(err)
	}
	latestData.SeasonData = temp.JsonContent.Season

	// Latest race from 2022 season:
	temp, err = f1api.GetFallBack("lastround_2022")
	if err != nil {
		log.Fatal(err)
	}
	latestData.RaceData = temp.JsonContent.Season.Races[0]

	// 2023 calendar data:
	temp, err = f1api.GetFallBack("calendar_2023")
	if err != nil {
		log.Fatal(err)
	}
	currentCalendar = temp.JsonContent.Season
	fmt.Println("RETRIEVED FALLBACK DATA SUCCESSFULLY.")
}
