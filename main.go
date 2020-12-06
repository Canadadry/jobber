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

	l.Main.Printf("duration %v", duration)
	if exitError, ok := err.(*exec.ExitError); ok {
		l.Main.Printf("ended with error : %v", exitError)

	} else {
		l.Main.Printf("ended succesfully")
	}

	return err
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
