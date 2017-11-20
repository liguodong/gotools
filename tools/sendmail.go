package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"gotools/util"
	"time"
)

const SEND_MAIL = "http://192.168.1.100:8199/v2/mails/send"

func SendMail (concurrency int, round int, sleep time.Duration) {
	if round == 0 {
		round++
	}
	for r := 1; r <= round; r++ {
		for i := 0; i < concurrency; i++ {
			go sendMail(i, r)
		}
		time.Sleep(sleep)
	}
}

func sendMail(seq int, round int) {
	postData := map[string]string {
		"to":		"liguodong@parllay.cn",
		"content":	fmt.Sprintf("ping [%d] pong..", seq),
		"subject":	fmt.Sprintf("mail service test round [%d], seq [%d]", round, seq),
	}

	headers := map[string]string {
		"Content-Type": 	"application/json",
		"Custom-Header": 	"charset=utf-8",
	}

	jsonByte, e := json.Marshal(postData)
	if e != nil {
		panic(fmt.Sprintf("error when marshal to a json: %s\n", e))
		return
	}

	if _, _, e := util.GetHttpResponse(SEND_MAIL, string(jsonByte), http.MethodPost, headers); e != nil {
		fmt.Printf("error when post to mail [%d] - [%d] service with reason: %v\n", round, seq, e)
		return
	} else {
		fmt.Printf("mail [%d] - [%d] sent!\n", round, seq)
	}
}