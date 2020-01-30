package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gotest.tools/golden"
	"os"
	"path/filepath"

	//"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()

	files, _ := filepath.Glob("testdata/*")
	for _, f := range files {
		if f != ".gitignore" {
			exclude, _ := filepath.Match("*/.gitignore", f)
			if !exclude {
				_ = os.Remove(f)
			}
		}
	}
}

func TestSuccessFullCopy(t *testing.T) {
	makeTmpFiles(t, "1234567890", "1234567890")
	err := copyTestFile(t, 0, 0)
	compareFiles(t)
	assert.NoError(t, err)
}

func TestSuccessLimitCopy(t *testing.T) {
	makeTmpFiles(t, "1234567890", "12345")
	err := copyTestFile(t, 5, 0)
	compareFiles(t)
	assert.NoError(t, err)
}

func TestSuccessOffsetCopy(t *testing.T) {
	makeTmpFiles(t, "1234567890", "67890")
	err := copyTestFile(t, 0, 5)
	compareFiles(t)
	assert.NoError(t, err)
}

func TestSuccessLimitAndOffsetCopy(t *testing.T) {
	makeTmpFiles(t, "1234567890", "67890")
	err := copyTestFile(t, 5, 5)
	compareFiles(t)
	assert.NoError(t, err)
}

func TestIncorrectOffsetCopy(t *testing.T) {
	err := copyTestFile(t, 0, 10)
	assert.Error(t, err)
}

func TestIncorrectLimitCopy(t *testing.T) {
	err := copyTestFile(t, -10, 0)
	assert.Error(t, err)
}

func TestIncorrectPath(t *testing.T) {
	err := copyTestFile(t, 0, 0)
	assert.Error(t, err)
}

func copyTestFile(t *testing.T, limit, offset int64) error {
	c := NewGoCopy()
	c.barEnabled = false

	return c.Copy(
		golden.Path(t.Name()+".input"),
		golden.Path(t.Name()+".output"),
		limit,
		offset,
	)
}

func makeFile(t *testing.T, value, postfix string) {
	inputFile, err := os.OpenFile(golden.Path(t.Name()+postfix), os.O_WRONLY|os.O_CREATE, 0755)

	if err != nil {
		panic(fmt.Sprintf("Cannot create test file. \n %#v", err))
	}

	defer inputFile.Close()

	_, _ = inputFile.WriteString(value)
}

func makeTmpFiles(t *testing.T, inputValue, expectedValue string) {
	makeFile(t, inputValue, ".input")
	makeFile(t, expectedValue, ".golden")
}

func compareFiles(t *testing.T) {
	golden.AssertBytes(t, golden.Get(t, t.Name()+".output"), t.Name()+".golden")
}
