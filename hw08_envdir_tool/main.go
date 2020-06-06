package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("arguments count is incorrect")
	}

	envDir := args[1]
	cmd := args[2:]

	env, err := ReadDir(envDir)
	if err != nil {
		log.Fatal(err)
	}

	exitCode := RunCmd(cmd, env)
	os.Exit(exitCode)
}
