package workflow

import (
	"github.com/liguodong/seelog"
	"fmt"
	"encoding/json"
	"strings"
)

type Executable interface {
	Execute(execution *Execution)
	Leave(execution *Execution)
}

/**
work flow model
 */
type Work_Flow struct {
	Id          int        `json:"id"`
	PlatId      string    `json:"platId"`
	BusinessId  int        `json:"businessId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Enable      bool    `json:"enable"`
	CreatedBy   int        `json:"createdBy"`
}

/**
flow element models
 */
type FlowElement struct {
	Id         int        `json:"id"`
	WorkFlowId int        `json:"workFlowId"`
	IsStart	   bool		  `json:"isStart"`
	Type       string    `json:"type"`
	Expression string    `json:"expression"`
	Incomings  []int    `json:"incomings"`
	Outgoings  []int    `json:"outgoings"`
}

type SequenceFlow struct {
	*FlowElement
}

func (p *SequenceFlow) Leave(execution *Execution) {
	seelog.Infof("Go to next element: %d", p.Outgoings[0])

	redisKey := fmt.Sprintf("redis_key_%d", p.Outgoings[0])

	flowElement, _ := loadFlowElementFromRedis(redisKey)

	switch flowElement.Type {
	case "trigger":
		nextElement := &TriggerTask{
			FlowElement: &flowElement,
		}
		execution.SetCurrentFlowElement(nextElement)
	case "exclusiveGateway":
		nextElement := &ExclusiveGateway{
			FlowElement: &flowElement,
		}
		execution.SetCurrentFlowElement(nextElement)
	case "task":
		nextElement := &Task{
			FlowElement: &flowElement,
		}
		execution.SetCurrentFlowElement(nextElement)
	}
}

func (p *SequenceFlow) Execute(execution *Execution) {
	seelog.Debugf("sequence...")
	p.Leave(execution)
}

type Task struct {
	*FlowElement
}

func (p *Task) Leave(execution *Execution) {
	if len(p.Outgoings) == 0 {
		seelog.Debugf("No outgoing element, work flow finished")
		return
	}

	seelog.Infof("Go to next element: %d", p.Outgoings[0])

	redisKey := fmt.Sprintf("redis_key_%d", p.Outgoings[0])

	flowElement, _ := loadFlowElementFromRedis(redisKey)

	sequence := &SequenceFlow{
		FlowElement: &flowElement,
	}

	execution.SetCurrentFlowElement(sequence)
}

func (p *Task) Execute(execution *Execution) {
	seelog.Debugf("task...")
	p.Leave(execution)
}

type TriggerTask struct {
	*FlowElement
}

func (p *TriggerTask) Leave(execution *Execution) {
	seelog.Infof("Go to next element: %d", p.Outgoings[0])

	redisKey := fmt.Sprintf("redis_key_%d", p.Outgoings[0])

	flowElement, _ := loadFlowElementFromRedis(redisKey)

	sequence := &SequenceFlow{
		FlowElement: &flowElement,
	}

	execution.SetCurrentFlowElement(sequence)
}

func (p *TriggerTask) Execute(execution *Execution) {
	seelog.Debugf("trigger...")
	p.Leave(execution)
}

type ExclusiveGateway struct {
	*FlowElement
}

func (p *ExclusiveGateway) Leave(execution *Execution) {
	seelog.Infof("next elements: %d, %d", p.Outgoings[0], p.Outgoings[1])

	for _, item := range p.Outgoings {
		redisKey := fmt.Sprintf("redis_key_%d",item)
		flowElement, _ := loadFlowElementFromRedis(redisKey)
		exp, _ := loadExpression(flowElement.Expression)
		sequence := &SequenceFlow{
			FlowElement: &flowElement,
		}
		if execution.Notification.Api == exp["api"] {
			execution.SetCurrentFlowElement(sequence)
		}
	}
}

func (p *ExclusiveGateway) Execute(execution *Execution) {
	seelog.Debugf("exclusive...")
	p.Leave(execution)
}

func loadFlowElementFromRedis(redisKey string) (flowElement FlowElement, err error) {
	err = json.NewDecoder(elements[redisKey]).Decode(&flowElement)
	if err != nil {
		seelog.Warnf("err: %s", err)
	}
	return
}

func loadExpression(exp string) (expression map[string]string, err error) {
	err = json.NewDecoder(strings.NewReader(exp)).Decode(&expression)
	if err != nil {
		seelog.Warnf("err: %s", err)
	}
	return
}
