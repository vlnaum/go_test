package main

import (
	"log"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		log.Fatal("arguments weren't defined")
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

	command := exec.Command(cmd[0], cmd[1:]...)
	command.Env = os.Environ()
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
		log.Fatal(err)
	}

	return
}
