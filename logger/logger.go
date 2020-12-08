package logger

import (
	"fmt"
	"github.com/canadadry/jobber/path"
	"io"
	"log"
	"os"
	"strings"
)

type Loggers struct {
	mainLog    io.Closer
	jobLogFile io.Closer
	Builder    *strings.Builder
	Main       *log.Logger
	Out        *log.Logger
	Sinker     *log.Logger
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

	l := Loggers{
		mainLog:    mainFile,
		jobLogFile: jobFile,
		Builder:    &strings.Builder{},
	}

	mainMux := io.MultiWriter(mainFile, l.Builder, os.Stdout)
	outMux := io.MultiWriter(jobFile, l.Builder, os.Stdout)

	l.Main = log.New(mainMux, p.JobId+" ", log.LstdFlags|log.Lmsgprefix)
	l.Out = log.New(outMux, p.JobId+"-out ", log.LstdFlags|log.Lmsgprefix)
	l.Sinker = log.New(outMux, p.JobId+"-sinker ", log.LstdFlags|log.Lmsgprefix)

	return l, nil
}

func (l Loggers) Close() {
	l.mainLog.Close()
	l.jobLogFile.Close()
}
