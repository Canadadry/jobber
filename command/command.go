package command

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func Runner(l *log.Logger, args ...string) func(string) (bool, error) {
	return func(command string) (bool, error) {
		if !exist(command) {
			return false, fmt.Errorf("%s does not exist", command)
		}

		cmd := exec.Command(command, args...)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			return false, err
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			return false, err
		}
		multi := io.MultiReader(stdout, stderr)

		if err := cmd.Start(); err != nil {
			return false, err
		}

		out := bufio.NewScanner(multi)
		scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			for i := 0; i < len(data); i++ {
				if data[i] == '\r' || data[i] == '\n' {
					return i + 1, data[:i], nil
				}
			}
			if !atEOF {
				return 0, nil, nil
			}
			return 0, data, bufio.ErrFinalToken
		})

		for out.Scan() {
			l.Printf("%s", out.Text())
		}
		if err := out.Err(); err != nil {
			l.Printf("error: %s", err)
		}

		err = cmd.Wait()
		_, ok := err.(*exec.ExitError)

		if err != nil && !ok {
			return false, fmt.Errorf("failed to execute %w", err)
		} else if err != nil {
			return false, nil
		}
		return true, nil
	}
}
