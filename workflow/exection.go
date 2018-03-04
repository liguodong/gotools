package workflow

import (
	"encoding/json"
	"github.com/liguodong/seelog"
	"github.com/pkg/errors"
)

type Execution struct {
	*Notification
}

func (p *Execution) SetCurrentFlowElement(element Executable) {
	element.Execute(p)
}

func (p *Execution) Verify(notification *Notification) (trigger TriggerTask, err error) {
	var flowElement FlowElement
	err = json.NewDecoder(elements["redis_key_1"]).Decode(&flowElement)
	if err != nil {
		seelog.Warnf("err: %s", err)
	}

	if flowElement.Type != "trigger" || !flowElement.IsStart {
		err = errors.New("Invalid trigger..")
	}

	trigger = TriggerTask{
		FlowElement: &flowElement,
	}
	return trigger, err
}