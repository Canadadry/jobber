package path

import (
	"fmt"
	"os"
)

const (
	jobFolder = "%s/.jobber/job/%s.sh"
	logFolder = "%s/.jobber/log/%s.log"
)

const (
	mainLogName = "history"
)

type Path struct {
	JobName string
	Job     string
	JobLog  string
	MainLog string
}

func Resolve(command string) (Path, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Path{}, err
	}

	return Path{
		JobName: command,
		Job:     fmt.Sprintf(jobFolder, home, command),
		JobLog:  fmt.Sprintf(logFolder, home, command),
		MainLog: fmt.Sprintf(logFolder, home, mainLogName),
	}, nil
}
