package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewList(t *testing.T) {
	list := NewList()

	assert.IsType(t, List{}, list)
}

// Тестируем добавление элемента в пустой список.
func TestList_PushFrontInEmptyList(t *testing.T) {
	list := NewList()
	assert.Equal(t, 0, list.Len())

	list.PushFront("1 element")
	assert.Equal(t, 1, list.Len())
	assert.Same(t, list.Last(), list.First())

	oldFirst := list.First()

	list.PushFront("2 element")
	list.PushFront("3 element")

	// Убеждаемся, что последний добавленный элемент в начало действительно идет первым.
	assert.Equal(t, "3 element", list.First().Value())

	// Смотрим что самый первый элемент сместился в конец.
	assert.Same(t, oldFirst, list.Last())
}

// Тестируем добавление элемента в список содержащий какие-то значения.
func TestList_PushFrontInFilledList(t *testing.T) {
	list := makeTestList()

	oldFirst := list.First()

	list.PushFront("first element")
	assert.Equal(t, "first element", list.First().Value())
	// Проверяем, что первый элемент сместился и стал вторым в списке
	assert.Same(t, oldFirst.Prev(), list.First())
	// Проверяем, что новый вставленный первый элемент ссылается на второй
	assert.Same(t, list.First().Next(), oldFirst)
}

// Тестируем порядок элементов в списке после добавления нескольких элементов в начало.
func TestList_PushFrontOrder(t *testing.T) {
	list := NewList()

	itemsCount := 10
	var i int

	for i = 0; i < itemsCount; i++ {
		list.PushFront(i)
	}

	item := list.First()

	for item != nil {
		i--
		assert.Equal(
			t,
			item.Value(),
			i,
			fmt.Sprintf("некорректный порядок элементов, ожидается элемент %d получен %d", i, item.Value()),
		)
		item = item.Next()
	}
}

func TestList_PushBackInFilledList(t *testing.T) {
	list := makeTestList()

	oldLast := list.Last()

	list.PushBack("last element")
	assert.Equal(t, "last element", list.Last().Value())

	// Проверяем, что последний элемент сместился и стал предпоследниим
	assert.Same(t, oldLast.Next(), list.Last())
	// Проверяем, что новый последний элемент ссылается на предыдущий
	assert.Same(t, list.Last().Prev(), oldLast)
}

// Тестируем добавление элемента в пустой список.
func TestList_PushBackInEmptyList(t *testing.T) {
	list := NewList()
	assert.Equal(t, 0, list.Len())

	list.PushBack("1 element")
	assert.Equal(t, 1, list.Len())
	assert.Equal(t, "1 element", list.Last().Value())
	assert.Same(t, list.Last(), list.First())

	oldLast := list.Last()

	list.PushBack("2 element")
	list.PushBack("3 element")

	// Убеждаемся, что последний добавленный элемент в конец находится в конце.
	assert.Equal(t, "3 element", list.Last().Value())
	// Смотрим что самый первый элемент остался вначале после добавления других элементов.
	assert.Same(t, oldLast, list.First())
}


// Тестируем порядок элементов в списке после добавления нескольких элементов в конец.
func TestList_PushBackOrder(t *testing.T) {
	list := NewList()

	itemsCount := 10
	var i int

	for i = 0; i < itemsCount; i++ {
		list.PushBack(i)
	}

	item := list.First()

	i = 0
	for item != nil {
		assert.Equal(
			t,
			item.Value(),
			i,
			fmt.Sprintf("некорректный порядок элементов, ожидается элемент %d получен %d", i, item.Value()),
		)
		i++
		item = item.Next()
	}
}
func TestList_First(t *testing.T) {
	list := NewList()
	assert.Nil(t, list.First())

	list.PushFront("test")
	assert.NotNil(t, list.First())
	assert.Equal(t,"test", list.First().Value())
}

func TestList_Last(t *testing.T) {
	list := NewList()
	assert.Nil(t, list.Last())

	list.PushBack("test")
	assert.NotNil(t, list.Last())
	assert.Equal(t,"test", list.Last().Value())
}

func TestList_Len(t *testing.T) {
	list := NewList()
	assert.Zero(t, list.Len())

	list.PushFront("test")
	assert.Equal(t, 1, list.Len())

	list.PushFront("test2")
	assert.Equal(t, 2, list.Len())
}

func TestList_Remove(t *testing.T) {
	list := makeTestList()

	initialLen := list.Len()
	// Берем второй элемент и дропаем его из списка.
	firstElement := list.First()
	secondElement := firstElement.Next()
	thirdElement := secondElement.Next()
	err := list.Remove(secondElement)

	assert.NoError(t, err)
	assert.Equal(t, initialLen - 1, list.Len())

	// Убеждаемся что теперь первый элемент ссылается на третий
	assert.Same(t, thirdElement, firstElement.Next())
}

func TestList_RemoveFirstElement(t *testing.T) {
	list := makeTestList()

	initialLen := list.Len()
	// Берем первый элемент из списка и удаляем его.
	firstElement := list.First()
	secondElement := firstElement.Next()
	err := list.Remove(firstElement)

	assert.Equal(t, initialLen - 1, list.Len())
	assert.NoError(t, err)

	// Убеждаемся что теперь второй элемент стал первым, и он больше не ссылается на предыдущий элемент.
	assert.Same(t, secondElement, list.First())
	assert.Nil(t, secondElement.Prev())
}

func TestList_RemoveLastElement(t *testing.T) {
	list := makeTestList()

	initialLen := list.Len()
	// Берем последний элемент из списка и удаляем его.
	lastElement := list.Last()
	prevLastElement := lastElement.Prev()

	err := list.Remove(lastElement)

	assert.Equal(t, initialLen - 1, list.Len())
	assert.NoError(t, err)

	// Убеждаемся что предпоследний элемент стал последним и не ссылается на следующий
	assert.Same(t, prevLastElement, list.Last())
	assert.Nil(t, prevLastElement.Next())
}

func makeTestList() List {
	itemsCount := 3 // В целом можно менять, но обязательно чтобы элементов было больше трех.

	list := NewList()
	for i := 0; i < itemsCount; i++ {
		list.PushFront(i)
	}

	return list
}