package api

import (
	"encoding/xml"
	"net/http"
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
