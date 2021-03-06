package main

import (
	"flag"
	"fmt"
	"github.com/canadadry/jobber/command"
	"github.com/canadadry/jobber/logger"
	"github.com/canadadry/jobber/path"
	"github.com/canadadry/jobber/timer"
	"time"
)

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

	var cmdSuccessfull bool
	var duration time.Duration

	runner := command.Runner(l.Out)
	duration = timer.TimeTask(func() {
		cmdSuccessfull, err = runner(path.Job)
	})

	if err != nil {
		l.Main.Printf("%v", err)
		return err
	}

	var sinkerEnv, sinkerPath string
	if cmdSuccessfull == false {
		l.Main.Printf("ended with error")
		sinkerEnv = env.Failure
		sinkerPath = path.Failure
	} else {
		l.Main.Printf("ended succesfully after %v", duration)
		sinkerEnv = env.Success
		sinkerPath = path.Success
	}

	if len(sinkerEnv) > 0 {
		l.Sinker.Printf("starting %s", sinkerEnv)
		_, err = command.Runner(l.Sinker, path.JobId, l.Builder.String())(sinkerPath)
		if err != nil {
			l.Sinker.Printf("ended with error %v", err)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Printf("error: %v", err)
	}
}
