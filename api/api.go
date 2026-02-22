package api

import (
	"encoding/xml"
	"net/http"
	"time"
	"vpdesktop/types"
)

func baseVPMobilRequest(url string, username string, password string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func VPMobilClassesRequest(url string, username string, password string) (types.ClassesResponse, error) {
	response, err := baseVPMobilRequest(url, username, password)
	if err != nil {
		return types.ClassesResponse{}, err
	}

	defer response.Body.Close()

	var result types.ClassesResponse

	err = xml.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func VPMobilTeachersRequest(url string, username string, password string) (types.TeachersResponse, error) {
	response, err := baseVPMobilRequest(url, username, password)
	if err != nil {
		return types.TeachersResponse{}, err
	}

	defer response.Body.Close()

	var result types.TeachersResponse

	err = xml.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil

}

func VPMobilRoomsRequest(url string, username string, password string) (types.RoomsResponse, error) {
	response, err := baseVPMobilRequest(url, username, password)
	if err != nil {
		return types.RoomsResponse{}, err
	}

	defer response.Body.Close()

	var result types.RoomsResponse

	err = xml.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil

}

func GetCurrentWeek(daysPerWeek int) (dateStart time.Time, dateEnd time.Time) {
	currentDate := time.Now()
	weekday := int(currentDate.Weekday())
	if weekday == 0 {
		weekday = 7
	}

	if weekday <= 5 {
		weekStart := currentDate.AddDate(0, 0, -weekday+1)
		weekEnd := weekStart.AddDate(0, 0, daysPerWeek-1)
		return weekStart, weekEnd
	} else {
		weekStart := currentDate.AddDate(0, 0, -weekday+8)
		weekEnd := weekStart.AddDate(0, 0, daysPerWeek-1)
		return weekStart, weekEnd
	}
}

func FetchWeeklyClasses(school, username, password string, daysPerWeek int) (types.WeeklyClassesResponse, error) {
	weekStart, weekEnd := GetCurrentWeek(daysPerWeek)

	var weeklyClasses types.WeeklyClassesResponse
	weeklyClasses.FetchStart = weekStart.Format("2006-01-02")
	weeklyClasses.FetchEnd = weekEnd.Format("2006-01-02")

	for date := weekStart; !date.After(weekEnd); date = date.AddDate(0, 0, 1) {
		url := ComposeURL("stundenplan24.de", types.PlanByClass, school, date, ".xml")
		response, err := VPMobilClassesRequest(url, username, password)

		if err != nil {
			return types.WeeklyClassesResponse{}, err
		}

		weeklyClasses.Classes = append(weeklyClasses.Classes, response)
	}
	return weeklyClasses, nil
}
