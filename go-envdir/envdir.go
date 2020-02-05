package main

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

// ReadDir function
func ReadDir(dir string) (map[string]string, error) {
	result := make(map[string]string)

	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return nil, err
	}

	err = checkPathIsDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		content, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, err
		}

		err = checkPathIsNotDir(f)
		if err != nil {
			return nil, err
		}

		result[filepath.Base(f)] = string(content)
	}

	return result, nil
}

func checkPathIsDir(path string) error {
	fInfo, err := os.Lstat(path)
	if err != nil {
		return err
	}

	if !fInfo.IsDir() {
		return errors.New("is not dir")
	}

	return nil
}

func checkPathIsNotDir(path string) error {
	fInfo, err := os.Lstat(path)
	if err != nil {
		return err
	}

	if fInfo.IsDir() {
		return errors.New("dir")
	}

	return nil
}

var outWriter io.Writer
var errWriter io.Writer

func RunCmd(cmd []string, env map[string]string) int {
	var err error
	for k, v := range env {
		if len(v) == 0 {
			err = os.Unsetenv(k)
		} else {
			err = os.Setenv(k, v)
		}

		if err != nil {
			return 111
		}
	}

	c := exec.Command(cmd[0], cmd[1:]...)
	if outWriter == nil {
		outWriter = os.Stdout
	}

	if errWriter == nil {
		errWriter = os.Stderr
	}

	c.Stdout = outWriter
	c.Stderr = errWriter

	err = c.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return exitErr.ExitCode()
		}

		return 111
	}

	return 0
}