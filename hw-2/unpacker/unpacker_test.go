package unpacker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStringData struct {
	Actual string
	Expected string
}

func TestUnpackIncorrectString(t *testing.T)  {
	u := New()

	var testData = []TestStringData {
		{
			Actual:   "1а",
		},
		{
			Actual:   "45",
		},
		{
			Actual:   `\`,
		},
		{
			Actual:   `\й`,
		},
	}

	for _, d := range testData {
		result, err := u.Unpack(d.Actual)

		assert.Empty(t, result, fmt.Sprintf("Actual: %s", d.Actual))
		assert.Error(t, err, fmt.Sprintf("Actual: %s", d.Actual))
	}
}

func TestUnpack(t *testing.T) {
	u := New()

	var testData = []TestStringData {
		{
			Actual:   "a1",
			Expected: "a",
		},
		{
			Actual:   "a2",
			Expected: "aa",
		},
		{
			Actual:   "a10",
			Expected: "aaaaaaaaaa",
		},
		{
			Actual:   "a10b",
			Expected: "aaaaaaaaaab",
		},
		{
			Actual:   "a1b",
			Expected: "ab",
		},
		{
			Actual:   "a4bc2d5e",
			Expected: "aaaabccddddde",
		},
		{
			Actual:   "abcd",
			Expected: "abcd",
		},
		{
			Actual:   `abcd\1`,
			Expected: "abcd1",
		},
		{
			Actual:   `a\\3`,
			Expected: `a\\\`,
		},
		{
			Actual:   `\\`,
			Expected: `\`,
		},
		{
			Actual:   `\\\\`,
			Expected: `\\`,
		},
		{
			Actual:   `qwe\45`,
			Expected: `qwe44444`,
		},
		{
			Actual:   `qwe\\5`,
			Expected: `qwe\\\\\`,
		},
		{
			Actual:   `qwe\210`,
			Expected: `qwe2222222222`,
		},
		{
			Actual:   `аббб\210`,
			Expected: `аббб2222222222`,
		},
		{
			Actual:   `а0`,
			Expected: ``,
		},
		{
			Actual:   `а00`,
			Expected: ``,
		},
		{
			Actual:   `а010`,
			Expected: `аааааааааа`,
		},
		{
			Actual:   `а2`,
			Expected: `аа`,
		},
	}

	for _, d := range testData {
		result, err := u.Unpack(d.Actual)

		assert.Equal(t, d.Expected, result, fmt.Sprintf("Unpack string: %s", d.Actual))
		assert.NoError(t, err, fmt.Sprintf("Unpack string: %s", d.Actual))
	}
}