package api

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"vpmobil_app/types"
)

func baseVPMobilRequest(urlType string, url string, username string, password string) (*http.Response, error) {

	fmt.Print(url)
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

	fmt.Println("\n" + res.Status)

	return res, nil
}

func VPMobilClassesRequest(url string, username string, password string) (types.ClassesResponse, error) {
	response, err := baseVPMobilRequest(types.Classes, url, username, password)
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
	response, err := baseVPMobilRequest(types.Teachers, url, username, password)
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
	response, err := baseVPMobilRequest(types.Rooms, url, username, password)
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
