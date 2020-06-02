package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "*")
	if err != nil {
		return
	}
	defer os.Remove(tempFile.Name())

	t.Run("positive case", func(t *testing.T) {
		result := Copy("testdata/test.txt", tempFile.Name(), 0, 0)
		tempFileStat, _ := tempFile.Stat()

		require.Nil(t, result)
		require.Equal(t, int32(10), int32(tempFileStat.Size()))
	})

	t.Run("offset exceeds file size", func(t *testing.T) {
		result := Copy("testdata/test.txt", tempFile.Name(), 11, 0)

		require.Equal(t, ErrOffsetExceedsFileSize, result)
	})

	t.Run("limit exceeds file size", func(t *testing.T) {
		result := Copy("testdata/test.txt", tempFile.Name(), 0, 100)
		tempFileStat, _ := tempFile.Stat()

		require.Nil(t, result)
		require.Equal(t, int32(10), int32(tempFileStat.Size()))
	})

	t.Run("unsupported file", func(t *testing.T) {
		result := Copy("testdata", tempFile.Name(), 0, 0)
		require.Equal(t, ErrUnsupportedFile, result)
	})
}
