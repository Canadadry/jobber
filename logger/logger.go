package logger

import (
	"fmt"
	"io"
	"jobber/path"
	"log"
	"os"
)

type Loggers struct {
	mainLog    io.Closer
	jobLogFile io.Closer
	Main       *log.Logger
	Out        *log.Logger
	Err        *log.Logger
}

func New(p path.Path) (Loggers, error) {

	mainFile, err := os.OpenFile(p.MainLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return Loggers{}, fmt.Errorf("cannot open main log file : %w", err)
	}
	jobFile, err := os.OpenFile(p.JobLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return Loggers{}, fmt.Errorf("cannot open job log file : %w", err)
	}

	return Loggers{
		mainLog:    mainFile,
		jobLogFile: jobFile,
		Main:       log.New(mainFile, p.JobName+" ", log.LstdFlags),
		Out:        log.New(jobFile, "out ", log.LstdFlags|log.Lmsgprefix),
		Err:        log.New(jobFile, "err ", log.LstdFlags|log.Lmsgprefix),
	}, nil
}

func (l Loggers) Close() {
	l.mainLog.Close()
	l.jobLogFile.Close()
}
