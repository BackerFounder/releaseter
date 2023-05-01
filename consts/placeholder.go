package consts

import (
	"test/config"
	"time"
)

func GetVersionPlaceholders() [3]string {
	return [3]string{"{{ $VERSION_NEXT_MAJOR }}", "{{ $VERSION_NEXT_MINOR }}", "{{ $VERSION_NEXT_PATCH }}"}
}

func GetTimePlaceholders() map[string]time.Time {
	return map[string]time.Time{
		"{{ $TIME_WORKFLOW }}": config.GetWorkflowTime(),
	}
}
