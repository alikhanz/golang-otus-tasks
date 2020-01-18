package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTop10(t *testing.T) {
	result := Top10("cat and ,dog one-two, - dog! two cats , and ?one man")
	assert.Subset(t, result, []string{"dog", "one", "and"})

	result = Top10(", -")
	assert.Nil(t, result)

	result = Top10("one-two")
	assert.Subset(t, result, []string{"one-two"})

	result = Top10("one one one one one one one one one one two two two two three three three four four five five six six seven seven eight eight nine nine ten ten eleven")
	assert.Subset(t, result, []string{"one", "two", "three"})
	assert.NotContains(t, result, "eleven")
}