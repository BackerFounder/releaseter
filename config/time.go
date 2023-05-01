package config

import "time"

var (
	workflowTime time.Time
)

func GetWorkflowTime() time.Time {
	if workflowTime.IsZero() {
		workflowTime = time.Now()
	}
	return workflowTime
}
