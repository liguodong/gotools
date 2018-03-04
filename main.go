package main

import (
	"gotools/tools"
	"gotools/workflow"
	"encoding/json"
	"strings"
	"github.com/liguodong/seelog"
)

func main() {
	seelog.Debugf("running")
	workflow.New()

	var str = strings.NewReader(`{
		"api": "url2",
		"platId": "L8B3Q17P47",
		"businessId": 170,
		"openId": "lk10r41j4gx1hxby3czxx6lh0000gn"
	}`)

	var notification workflow.Notification

	err := json.NewDecoder(str).Decode(&notification)
	if err != nil {
		seelog.Warnf("err: %s", err)
	}

	tools.HandleNotification(&notification)
	for true {}
}