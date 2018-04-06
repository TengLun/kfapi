package kfapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest(accountID, view, startDate, endDate, format, fraudType, endpoint, authKey string, filters []filter) (KFResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var req request
	req.AccountID = accountID
	req.View = "account"
	req.StartDate = startDate
	req.EndDate = endDate
	req.Format = format
	req.FraudType = fraudType
	req.Filters = filters

	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return KFResponse{}, err
	}

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return KFResponse{}, err
	}

	fmt.Printf("%#v\n", request)
	fmt.Println(string(reqBody))
	request.Header.Add("Authentication-Key", authKey)
	machine := &http.Client{}

	res, err := machine.Do(request)
	if err != nil {
		fmt.Println(err)
		return KFResponse{}, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return KFResponse{}, err
	}

	fmt.Println(string(resBody))

	var resp KFResponse

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		fmt.Println(err)
		return KFResponse{}, err
	}

	return resp, nil

}
