package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/canadadry/jobber/command"
	"github.com/canadadry/jobber/logger"
	"github.com/canadadry/jobber/path"
	"github.com/canadadry/jobber/timer"
	"log"
)

func ScanAndLog(b *bytes.Buffer, l *log.Logger) {
	in := bufio.NewScanner(b)
	for in.Scan() {
		l.Printf(in.Text())
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}
}

func run() error {

	var env path.Env
	flag.StringVar(&env.Failure, "f", "", "define the sinker to run on failure")
	flag.StringVar(&env.Success, "s", "", "define the sinker to run on successs")
	flag.StringVar(&env.Command, "j", "", "define the job to run")
	flag.Parse()

	if len(env.Command) == 0 {
		return fmt.Errorf("argument j cannot be empty")
	}

	path, err := path.Resolve(env)
	if err != nil {
		return fmt.Errorf("cannot find home dir : %w", err)
	}

	l, err := logger.New(path)
	if err != nil {
		return err
	}
	defer l.Close()

	var ok bool
	var stdout, stderr bytes.Buffer

	runner := command.Runner(&stdout, &stderr)
	duration := timer.TimeTask(func() {
		ok, err = runner(path.Job)
	})

	ScanAndLog(&stdout, l.Out)
	ScanAndLog(&stderr, l.Err)

	if err != nil {
		l.Main.Printf("%v", err)
		return err
	}
	if !ok {
		l.Main.Printf("ended with error")
		if len(env.Failure) > 0 {
			l.Main.Printf("starting sinker %s", env.Failure)
			var stdout, stderr bytes.Buffer
			runner := command.Runner(&stdout, &stderr)
			_, err = runner(path.Failure)
			if err != nil {
				l.Err.Printf("sinker ended with error %v", err)
			}

			ScanAndLog(&stdout, l.Sinker)
			ScanAndLog(&stderr, l.Sinker)
		}
	} else {
		l.Main.Printf("ended succesfully after %v", duration)
		if len(env.Success) > 0 {
			l.Main.Printf("starting sinker %s", env.Success)
			var stdout, stderr bytes.Buffer
			runner := command.Runner(&stdout, &stderr)
			_, err = runner(path.Success)
			if err != nil {
				l.Err.Printf("sinker ended with error %v", err)
			}
			ScanAndLog(&stdout, l.Sinker)
			ScanAndLog(&stderr, l.Sinker)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %v", err)
	}
}
