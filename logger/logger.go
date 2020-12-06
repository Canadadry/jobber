package logger

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/canadadry/jobber/path"
	"io"
	"log"
	"os"
)

type Loggers struct {
	mainLog    io.Closer
	jobLogFile io.Closer
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

	return Loggers{
		mainLog:    mainFile,
		jobLogFile: jobFile,
		Main:       log.New(mainFile, p.JobId+" ", log.LstdFlags|log.Lmsgprefix),
		Out:        log.New(jobFile, p.JobId+"-out ", log.LstdFlags|log.Lmsgprefix),
		Sinker:     log.New(jobFile, p.JobId+"-sinker ", log.LstdFlags|log.Lmsgprefix),
	}, nil
}

func (l Loggers) Close() {
	l.mainLog.Close()
	l.jobLogFile.Close()
}

func CaptureInLog(bout *bytes.Buffer, lout *log.Logger, task func(w io.Writer)) {
	if bout == nil {
		bout = &bytes.Buffer{}
	}
	task(bout)
	ScanAndLog(bout, lout)
}

func ScanAndLog(b *bytes.Buffer, l *log.Logger) {
	in := bufio.NewScanner(b)
	for in.Scan() {
		l.Printf(in.Text())
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}
}
