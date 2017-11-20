package util

import (
	"time"
	"net/http"

	"github.com/parnurzeal/gorequest"
	"github.com/liguodong/seelog"
	"github.com/go-errors/errors"
	"fmt"
)

func GetHttpResponse(url string, data string, method string, headers ...map[string]string) (gorequest.Response, string, error) {
	var client = gorequest.New().Timeout(5 * time.Second)

	switch method {
	case http.MethodPost:
		agent := client.Post(url)
		if headers != nil {
			agent.Header = headers[0]
		}

		seelog.Infof("data to post: %s", data)

		resp, body, errs := agent.Send(data).End(printResponse)
		if errs != nil {
			destroy(client)
			err := errors.New(errs)
			return nil, "", err
		}
		if resp.StatusCode > 399 {
			destroy(client)
			err := errors.New(fmt.Sprintf("bad request with code [%d], message: %s", resp.StatusCode, body))
			return nil, "", err
		}
		destroy(client)
		return resp, body, nil
	case http.MethodGet:
		agent := client.Get(url)
		if headers != nil {
			agent.Header = headers[0]
		}

		resp, body, errs := agent.End(printResponse)
		if errs != nil {
			destroy(client)
			err := errors.New(errs)
			return nil, "", err
		}
		if resp.StatusCode > 399 {
			destroy(client)
			err := errors.New(fmt.Sprintf("bad request with code [%d], message: %s", resp.StatusCode, body))
			return nil, "", err
		}
		destroy(client)
		return resp, body, nil
	}

	destroy(client)
	return nil, "", nil
}

func destroy(client *gorequest.SuperAgent)  {
	client = nil
}

func printResponse(resp gorequest.Response, body string, errs []error) {
	if errs != nil {
		seelog.Warnf("request to: %s\n<== some errors happened during remote call: %s", resp.Request.URL.String(), errs)
	}
	seelog.Infof("request to: %s\n<== finished with code: %s, value: %s", resp.Request.URL.String(), resp.Status, body)
}