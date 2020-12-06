package main

import (
	"fmt"
	"jobber/command"
	"jobber/path"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	LogPath = "%s/.jobber/log/%s.log"
)

func timeTask(task func()) time.Duration {
	start := time.Now()

	task()

	return time.Since(start)
}

func run() error {
	if len(os.Args) <= 1 {
		return fmt.Errorf("usage : jobber jobname")
	}
	path, err := path.Resolve(os.Args[1])
	if err != nil {
		return fmt.Errorf("cannot find home dir : %w", err)
	}

	f, err := os.OpenFile(path.MainLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("cannot open main log file : %w", err)
	}
	defer f.Close()

	logger := log.New(f, os.Args[1]+" ", log.LstdFlags)
	logger.Println("started")
	runner := command.Runner(os.Stdout, os.Stderr)
	duration := timeTask(func() {
		err = runner(path.Job)
	})
	logger.Printf("duration %v", duration)
	if exitError, ok := err.(*exec.ExitError); ok {
		logger.Printf("ended with error : %v", exitError)

	} else {
		logger.Printf("ended succesfully")
	}
	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
