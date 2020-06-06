package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("execution of command without args returns error exit code", func(t *testing.T) {
		exitCode := RunCmd(nil, nil)
		require.Equal(t, 1, exitCode)
	})

	t.Run("successfull execution of command returns exit code", func(t *testing.T) {
		command := []string{"bash", "exit 5"}
		result := RunCmd(command, nil)
		expectedResult := 5

		require.Equal(t, expectedResult, result)
	})

	t.Run("env vars", func(t *testing.T) {
		os.Setenv("VAR1", "test1")
		os.Setenv("VAR2", "test2")

		env := Environment{
			"VAR0": "test0",
			"VAR1": "",
			"VAR2": "testValue",
		}

		command := []string{"bash"}

		exitCode := RunCmd(command, env)
		require.Equal(t, 0, exitCode)

		var0, ok := os.LookupEnv("VAR0")
		require.True(t, ok)
		require.Equal(t, "test0", var0)

		_, ok := os.LookupEnv("VAR1")
		require.False(t, ok)

		var2, ok := os.LookupEnv("VAR2")
		require.True(t, ok)
		require.Equal(t, "testValue", var2)
	})
}
