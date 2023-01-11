package f1api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func (container JsonContainer) isEmpty() bool {
	return (len(container.JsonContent.Season.Races)) == 0
}

var months = []string{"", "JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG", "SEPT", "OCT", "NOV", "DEC"}

func APIresponse(url string) (response JsonContainer, err error) {

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	res, err := client.Do(req)
	if err != nil {
		return response, fmt.Errorf("error when trying to reach api with URL '%v'", url)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return response, fmt.Errorf("error request response with status code %d", res.StatusCode)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, fmt.Errorf("error when trying to read api response with URL '%v", url)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("error when trying to unmarshal api response with url '%v'\nERROR: '%v'", url, err)
	}

	if response.isEmpty() {
		return response, fmt.Errorf("error: after unmarshaling, the struct JsonContainer seems empty\nIf URL '%v' is correct, the api might be unreachable at the moment", url)
	}
	return response, nil
}

func GetFallBack(filename string) (temp JsonContainer, err error) {
	data, err := os.ReadFile(fmt.Sprintf("./data/%v.json", filename))
	if err != nil {
		log.Fatal("Error when opening file: ", err)

	}
	err = json.Unmarshal(data, &temp)
	return temp, err
}

func (race Race) GetTimeInterval() map[string][]string {
	interval := make(map[string][]string)
	startingDate := DateToArr(race.Qualifying.Date)
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

func FormatTime(date string) string {
	result := strings.Split(date, "-")
	temp, _ := strconv.Atoi(result[1])
	result[1] = months[temp]
	result[0], result[2] = result[2], result[0]

	return strings.Join(result, " ")
}
