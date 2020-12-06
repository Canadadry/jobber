package main

import (
	"bufio"
	"bytes"
	"fmt"
	"jobber/command"
	"jobber/logger"
	"jobber/path"
	"jobber/timer"
	"log"
	"os"
	"os/exec"
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
	if len(os.Args) <= 1 {
		return fmt.Errorf("usage : jobber jobname")
	}
	commandName := os.Args[1]
	path, err := path.Resolve(commandName)
	if err != nil {
		return fmt.Errorf("cannot find home dir : %w", err)
	}

	l, err := logger.New(path)
	if err != nil {
		return err
	}
	defer l.Close()

	var stdout, stderr bytes.Buffer

	runner := command.Runner(&stdout, &stderr)
	duration := timer.TimeTask(func() {
		err = runner(path.Job)
	})

	ScanAndLog(&stdout, l.Out)
	ScanAndLog(&stderr, l.Err)

	_, ok := err.(*exec.ExitError)

	if err != nil && !ok {
		err = fmt.Errorf("failed to execute %w", err)
		l.Main.Printf("%v", err)
		return err
	} else if err != nil {
		l.Main.Printf("ended with error : %v", err)
	} else {
		l.Main.Printf("ended succesfully after %v", duration)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
