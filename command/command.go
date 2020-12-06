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

func Runner(stdout io.Writer, stderr io.Writer) func(string) error {
	return func(command string) error {
		if !exist(command) {
			return fmt.Errorf("%s does not exist", command)
		}
		cmd := exec.Command(command)
		cmd.Stdout = stdout
		cmd.Stderr = stderr
		return cmd.Run()
	}
}
