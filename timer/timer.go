package timer

import (
	"time"
)

func TimeTask(task func()) time.Duration {
	start := time.Now()

	task()

	return time.Since(start)
}
