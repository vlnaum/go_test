package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]string

func ReadDir(dir string) (Environment, error) {
	dirs, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)

	for _, fileInfo := range dirs {
		fileName := fileInfo.Name()
		fileFullName := filepath.Join(dir, fileName)

		content, err := getValueFromFile(fileFullName)
		if err != nil {
			return nil, err
		}

		env[fileName] = content
	}

	return env, nil
}

func getValueFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	content, _, err := reader.ReadLine()
	if err != nil {
		return "", err
	}

	content = bytes.ReplaceAll(content, []byte("\x00"), []byte("\n"))

	return strings.TrimRight(string(content), "\n\t "), nil
}
