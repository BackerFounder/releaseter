package base

import (
	"time"
)

var (
	workflowTime = time.Now()
)

func GetWorkflowTime() time.Time {
	return workflowTime
}
