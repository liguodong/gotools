package tools

import (
	"gotools/workflow"
	"github.com/liguodong/seelog"
)

func HandleNotification(notification *workflow.Notification)  {
	seelog.Debugf("Suppose we receive a notify from service: %s, now generate an execution...", notification.Api)
	execution := &workflow.Execution{
		Notification: notification,
	}
	trigger, err := execution.Verify(notification)

	if err != nil {
		panic(err)
	}

	execution.SetCurrentFlowElement(&trigger)
}