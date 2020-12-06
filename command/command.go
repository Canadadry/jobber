package command

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

const (
	JobPath    = "%s/.jobber/job/%s.sh"
	SinkerPath = "%s/.jobber/sinker/%s.sh"
)

func exist(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

func Runner(path string, stdout, stderr io.Writer) func(string) error {
	return func(command string) error {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		command = fmt.Sprintf(path, home, command)
		if !exist(command) {
			return fmt.Errorf("%s does not exist", command)
		}
		cmd := exec.Command(command)
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		return cmd.Run()
	}
}
