package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "*")
	if err != nil {
		return
	}
	defer os.RemoveAll(tempDir)

	t.Run("positive case", func(t *testing.T) {
		env, err := ReadDir("testdata/env")
		require.NoError(t, err)

		expected := Environment{
			"BAR":   "bar",
			"FOO":   "   foo\nwith new line",
			"HELLO": "\"hello\"",
			"UNSET": "",
		}

		require.Equal(t, expected, env)
	})

	t.Run("= in env name", func(t *testing.T) {
		filePath := filepath.Join(tempDir, "=TEST=ENV=")
		err := ioutil.WriteFile(filePath, []byte("test"), 0777)
		require.NoError(t, err)

		defer os.Remove(filePath)

		env, err := ReadDir(tempDir)
		require.NoError(t, err)

		expected := Environment{
			"TESTENV": "test",
		}

		require.Equal(t, expected, env)
	})

	t.Run("empty directory", func(t *testing.T) {
		env, err := ReadDir(tempDir)
		require.NoError(t, err)

		expected := Environment{}

		require.Equal(t, expected, env)
	})

	t.Run("directory is not exist", func(t *testing.T) {
		env, err := ReadDir("testdata/dir")
		require.Nil(t, env)
		require.Error(t, err)
	})
}
