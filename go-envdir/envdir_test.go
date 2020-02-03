package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

const testDataDir = "testdata"

func TestMain(m *testing.M) {
	// Папка может быть не удалена, если какой-то из тестов сфейлился
	_ = os.RemoveAll(testDataDir)
	err := os.Mkdir(testDataDir, 0755)

	if err != nil {
		panic(err)
	}

	m.Run()

	// переживем если папка не удалится
	_ = os.RemoveAll(testDataDir)
}

func TestReadDir(t *testing.T) {
	var envs = map[string]string{
		"SOME_VAR": "SOME_VAL",
		"EMPTY_VAR": "",
	}

	err := os.Setenv("EMPTY_VAR", "SOME_TEST_DATA")
	if err != nil {
		t.Fatalf("setenv failed %#v", err)
	}

	makeTestEnvs(t, envs)
	resEnvs, err := ReadDir(testDir(t))

	assert.NoError(t, err)

	for k, v := range envs {
		assert.Contains(t, resEnvs, k)
		assert.Equal(t, resEnvs[k], v)
	}
}

func TestReadNotExistDir(t *testing.T) {
	_, err := ReadDir(testDir(t))
	assert.Error(t, err)
}

func TestRunCmd(t *testing.T) {
	var envs = map[string]string{
		"SOME_VAR": "SOME_VAL",
		"EMPTY_VAR": "",
	}

	w := &bytes.Buffer{}

	outWriter = w
	errWriter = w
	RunCmd([]string{"env"}, envs)
	assert.Contains(t, w.String(), "SOME_VAR=SOME_VAL")
	assert.NotContains(t, w.String(), "EMPTY_VAR")
}

func TestRunErrorCmd(t *testing.T) {
	w := &bytes.Buffer{}

	outWriter = w
	errWriter = w
	code := RunCmd([]string{"any-not-known-programm"}, map[string]string{})

	assert.Equal(t, code, 111)
}

func makeTestEnvs(t *testing.T, envs map[string]string) {
	dir := testDir(t)
	err := os.Mkdir(dir, 0755)

	if err != nil {
		t.Fatalf("cannot create dir %s err: \n %#v", dir, err)
	}

	for k, v := range envs {
		fPath := filepath.Join(dir, k)
		f, err := os.Create(fPath)

		if err != nil {
			t.Fatalf("cannot create file %s err: \n %#v", fPath, err)
		}

		_, err = f.WriteString(v)

		if err != nil {
			t.Fatalf("failed writing to file %s err: \n %#v", fPath, err)
		}

		f.Close()
	}
}

func testDir(t *testing.T) string {
	return filepath.Join(testDataDir, t.Name())
}