package main

import (
	"fmt"
	"time"

	"github.com/tenglun/kfapi"
)

func main() {
	acc := kfapi.Account{
		AccountID: "4225",
		ApiKey:    "3AC122C0-F124-59F1-604F-F7EE86038BBA",
		Type:      "account",
	}

	err := acc.ListApps(kfapi.PlatformDifference, time.Unix(1510000000, 0), time.Now())
	if err != nil {
		fmt.Println(err)
	}
}
