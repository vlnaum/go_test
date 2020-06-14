package main

import (
	"os"
	"os/exec"
)

const (
	exitCodeSuccess = 0
	exitCodeFailure = 1
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return exitCodeFailure
	}

	for k, v := range env {
		if v == "" {
			os.Unsetenv(k)
		} else {
			_, ok := os.LookupEnv(k)
			if ok {
				os.Unsetenv(k)
			}
			os.Setenv(k, v)
		}
	}

	command := exec.Command(cmd[0], cmd[1:]...) //nolint:gosec
	command.Env = os.Environ()
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
		return exitCodeFailure
	}

	return exitCodeSuccess
}
