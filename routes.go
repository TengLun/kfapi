package kfapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func sendRequest(r request) (fraudresponse, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	reqBody, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return fraudresponse{}, err
	}

	req, err := http.NewRequest("POST", "https://fraud.api.kochava.com:8320/fraud/platformdiff/network/data", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return fraudresponse{}, err
	}

	machine := &http.Client{}

	res, err := machine.Do(req)
	if err != nil {
		return fraudresponse{}, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fraudresponse{}, err
	}

	var resp fraudresponse

	err = json.Unmarshal(resBody, resp)
	if err != nil {
		return fraudresponse{}, err
	}

	return resp, nil
}
