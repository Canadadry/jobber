package main

import (
	"fmt"
	"jobber/command"
	"os"
)

func run() error {
	if len(os.Args) <= 1 {
		return fmt.Errorf("usage : jobber jobname")
	}
	runner := command.Runner(command.JobPath, os.Stdout, os.Stderr)
	return runner(os.Args[1])
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
