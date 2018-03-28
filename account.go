package kfapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Account struct {
	AccountID string
	Format    string
	Type      string
	ApiKey    string
}

func (a Account) ListApps(f string, startDate time.Time, endDate time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[f] + `/list/apps`

	var req request
	req.AccountID = a.AccountID
	req.View = a.Type
	req.StartDate = startDate.Format("2006-1-2")
	req.EndDate = endDate.Format("2006-1-2")
	req.Format = "json"
	req.FraudType = f

	reqBody, err := json.Marshal(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println(err)
		return err
	}

	request.Header.Add("Authentication-Key", a.ApiKey)
	machine := &http.Client{}

	res, err := machine.Do(request)
	if err != nil {
		fmt.Println(err)
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var resp fraudresponse

	err = json.Unmarshal(resBody, &resp)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(resBody))
	fmt.Printf("%#v\n", resp)

	return nil
}

func (a Account) ListNetworks(f string, startDate time.Time, endDate time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	return nil
}

func (a Account) ListAccounts(f string, startDate time.Time, endDate time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	return nil
}

func (a Account) FullData(f string, startDate time.Time, endDate time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	return nil
}

func (a Account) AccountsData(f string, startDate time.Time, endDate time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	return nil
}

func (a Account) AppData(f string, startDate time.Time, endDate time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	return nil
}

func (a Account) SiteIdData(f string, startDate time.Time, endDate time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	return nil
}

func (a Account) TrackerData(f string, startDate time.Time, endDate time.Time) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)

		}
	}()
	return nil
}
