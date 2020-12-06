package command

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func exist(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	f.Close()
	return true
}

func Runner(stdout io.Writer, stderr io.Writer) func(string) (bool, error) {
	return func(command string) (bool, error) {
		if !exist(command) {
			return false, fmt.Errorf("%s does not exist", command)
		}
		cmd := exec.Command(command)
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		err := cmd.Run()

		_, ok := err.(*exec.ExitError)

		if err != nil && !ok {
			return false, fmt.Errorf("failed to execute %w", err)
		} else if err != nil {
			return false, nil
		}
		return true, nil
	}
}
