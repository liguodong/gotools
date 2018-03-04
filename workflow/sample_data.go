package workflow

import "strings"

var elements = map[string]*strings.Reader{
	"redis_key_1": ELEMENT_A,
	"redis_key_2": ELEMENT_B,
	"redis_key_3": ELEMENT_C,
	"redis_key_4": ELEMENT_D,
	"redis_key_5": ELEMENT_E,
	"redis_key_6": ELEMENT_F,
	"redis_key_7": ELEMENT_G,
	"redis_key_8": ELEMENT_H,
	"redis_key_9": ELEMENT_I,
	"redis_key_10": ELEMENT_J,
	"redis_key_11": ELEMENT_K,
}

var (
	WORK_FLOW_1 = strings.NewReader(`{
		"id": 1,
		"platId": "L8B3Q17P47",
		"businessId": 170,
		"title": "work flow 1",
		"description": "this is a work flow...",
		"enable": true,
		"createdBy": 1
	}`)

	WORK_FLOW_2 = strings.NewReader(`{
		"id": 2,
		"platId": "L8B3Q17P47",
		"businessId": 170,
		"title": "work flow 2",
		"description": "this is a work flow...",
		"enable": true,
		"createdBy": 1
	}`)

	ELEMENT_A = strings.NewReader(`{
		"id": 1,
		"workFlowId": 1,
		"isStart": true,
		"type": "trigger",
		"incomings": [],
		"outgoings": [2]
	}`)

	ELEMENT_B = strings.NewReader(`{
		"id": 2,
		"workFlowId": 1,
		"type": "sequence",
		"incomings": [1],
		"outgoings": [3]
	}`)

	ELEMENT_C = strings.NewReader(`{
		"id": 3,
		"workFlowId": 1,
		"type": "exclusiveGateway",
		"incomings": [2],
		"outgoings": [4, 5]
	}`)

	ELEMENT_D = strings.NewReader(`{
		"id": 4,
		"workFlowId": 1,
		"type": "sequence",
		"expression": "{\"api\":\"url1\"}",
		"incomings": [3],
		"outgoings": [6]
	}`)

	ELEMENT_E = strings.NewReader(`{
		"id": 5,
		"workFlowId": 1,
		"type": "sequence",
		"expression": "{\"api\":\"url2\"}",
		"incomings": [3],
		"outgoings": [7]
	}`)

	ELEMENT_F = strings.NewReader(`{
		"id": 6,
		"workFlowId": 1,
		"type": "task",
		"incomings": [4],
		"outgoings": []
	}`)

	ELEMENT_G = strings.NewReader(`{
		"id": 7,
		"workFlowId": 1,
		"type": "task",
		"incomings": [5],
		"outgoings": [8]
	}`)

	ELEMENT_H = strings.NewReader(`{
		"id": 8,
		"workFlowId": 1,
		"type": "sequence",
		"expression": "",
		"incomings": [7],
		"outgoings": [9]
	}`)

	ELEMENT_I = strings.NewReader(`{
		"id": 9,
		"workFlowId": 1,
		"type": "trigger",
		"expression": "",
		"incomings": [8],
		"outgoings": [10]
	}`)

	ELEMENT_J = strings.NewReader(`{
		"id": 10,
		"workFlowId": 1,
		"type": "sequence",
		"expression": "",
		"incomings": [9],
		"outgoings": [11]
	}`)

	ELEMENT_K = strings.NewReader(`{
		"id": 11,
		"workFlowId": 1,
		"type": "task",
		"incomings": [10],
		"outgoings": []
	}`)
)
