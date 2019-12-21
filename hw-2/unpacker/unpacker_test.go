package unpacker

import (
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
	}

	for _, d := range testData {
		result, err := u.Unpack(d.Actual)

		assert.Empty(t, result)
		assert.Error(t, err)
	}
}

func TestUnpack(t *testing.T) {
	u := New()

	var testData = []TestStringData {
		{
			Actual:   "а1",
			Expected: "а",
		},
		{
			Actual:   "а10",
			Expected: "аааааааааа",
		},
		{
			Actual:   "а10б",
			Expected: "ааааааааааб",
		},
		{
			Actual:   "а1б",
			Expected: "аб",
		},
		{
			Actual:   "a4bc2d5e",
			Expected: "aaaabccddddde",
		},
		{
			Actual:   "abcd",
			Expected: "abcd",
		},
	}

	for _, d := range testData {
		result, err := u.Unpack(d.Actual)

		assert.Equal(t, d.Expected, result)
		assert.NoError(t, err)
	}
}