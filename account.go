package kfapi

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Account struct {
	List           List
	GatherDataFrom GatherDataFrom
}

type List struct {
	AccountID string
	Format    string
	View      string
	ApiKey    string
}

type GatherDataFrom struct {
	AccountID string
	Format    string
	View      string
	ApiKey    string
}

func GetAccount(apiKey string, accountID string) (*Account, error) {

	var account Account

	view, err := getView(apiKey)

	if err != nil {
		fmt.Println(err)
		return &Account{}, err
	}

	account.List.AccountID = accountID
	account.List.ApiKey = apiKey
	account.List.View = view

	account.GatherDataFrom.AccountID = accountID
	account.GatherDataFrom.ApiKey = apiKey
	account.GatherDataFrom.View = view

	return &account, nil
}

func getView(apiKey string) (string, error) {
	endpoint := "https://fraud.api.kochava.com:8320/fraud/installreceipt/tracker/data"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer([]byte(`
{
  "view": "network",
  "fraudType": "installReceiptVerification",
  "accountId": "XXX",
  "startDate": "2016-11-13",
  "endDate": "2017-1-11",
  "format": "JSON",
  "filters": []
}`)))

	machine := &http.Client{}

	res, err := machine.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	switch res.StatusCode {
	case 200:
		return "network", nil
	case 403:
		return "account", nil
	default:
		return "", errors.New(res.Status)
	}

}

func (l List) Apps(fraudType string, startDate, endDate time.Time, filters ...filter) (KFResponse, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/list/apps`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.ApiKey, filters)

}

func (l List) Networks(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (KFResponse, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/list/networks`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.ApiKey, filters)

}

func (l List) Accounts(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (KFResponse, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/list/accounts`

	return sendRequest(l.AccountID, l.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, l.ApiKey, filters)

}

func (g GatherDataFrom) Accounts(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (KFResponse, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.ApiKey, filters)

}

func (g GatherDataFrom) Apps(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (KFResponse, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/app/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.ApiKey, filters)

}

func (g GatherDataFrom) SiteIds(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (KFResponse, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/siteid/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.ApiKey, filters)

}

func (g GatherDataFrom) Trackers(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (KFResponse, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/tracker/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.ApiKey, filters)

}

func (g GatherDataFrom) Networks(fraudType string, startDate time.Time, endDate time.Time, filters ...filter) (KFResponse, error) {

	endpoint := `https://fraud.api.kochava.com:8320/fraud/` + fraudEndpointMap[fraudType] + `/network/data`

	return sendRequest(g.AccountID, g.View, startDate.Format("2006-1-2"), endDate.Format("2006-1-2"), "json", fraudType, endpoint, g.ApiKey, filters)

}

type KFResponse struct {
	MetaData struct {
		Headers []string `json:"headers"`
	} `json:"metaData"`
	Data []struct {
		AppName         string `json:"appName,omitempty"`
		AppID           string `json:"appId,omitempty"`
		NetworkName     string `json:"networkName,omitempty"`
		NetworkID       string `json:"networkId,omitempty"`
		ClickCt         int    `json:"clickCt,omitempty"`
		SameAcctClickCt int    `json:"sameAcctClickCt,omitempty"`
		DiffAcctClickCt int    `json:"diffAcctClickCt,omitempty"`
		InstallCt       int    `json:"installCt,omitempty"`
	} `json:"data"`
}

// SetThreshold returns an object with only the offending site_ids, IPs, and devices.
// Use it to set the threshold you consider unacceptable.
/*
func (k *KFResponse) SetThreshold(metric string, comparator string, threshold int) error {

	for i := range k.Data {
		v := *k.FieldByName(metric)
	}

	return k, nil
}
*/
type Threshold struct {
	AppName         string
	AppID           string
	NetworkName     string
	NetworkID       string
	ClickCt         int
	SameAcctClickCt int
	DiffAcctClickCt int
	InstallCt       int
}
