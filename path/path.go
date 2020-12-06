package path

import (
	"fmt"
	"os"
	"time"
)

const (
	jobFolder    = "%s/.jobber/job/%s.sh"
	sinkerFolder = "%s/.jobber/sinker/%s.sh"
	logFolder    = "%s/.jobber/log/%s.log"
)

const (
	mainLogName = "history"
)

type Path struct {
	JobId   string
	Job     string
	Failure string
	Success string
	JobLog  string
	MainLog string
}

type Env struct {
	Command string
	Success string
	Failure string
}

func Resolve(e Env) (Path, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Path{}, err
	}

	return Path{
		JobId:   fmt.Sprintf("%s-%d", e.Command, time.Now().Unix()),
		Job:     fmt.Sprintf(jobFolder, home, e.Command),
		Failure: fmt.Sprintf(sinkerFolder, home, e.Failure),
		Success: fmt.Sprintf(sinkerFolder, home, e.Success),
		JobLog:  fmt.Sprintf(logFolder, home, e.Command),
		MainLog: fmt.Sprintf(logFolder, home, mainLogName),
	}, nil
}
