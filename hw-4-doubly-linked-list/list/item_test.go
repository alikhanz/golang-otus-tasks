package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewItem(t *testing.T) {
	item := NewItem(``)

	assert.IsType(t, Item{}, item)
}

func TestItem_Next(t *testing.T) {
	item := NewItem(`1`)
	item2 := NewItem(`2`)

	item.next = &item2

	assert.Same(t, &item2, item.Next())
}

func TestItem_Prev(t *testing.T) {
	item := NewItem(``)
	item2 := NewItem(``)

	item.prev = &item2

	assert.Same(t, &item2, item.Prev())
}

func TestItem_Value(t *testing.T) {
	item := NewItem(`test`)

	assert.Equal(t, `test`, item.Value())
}