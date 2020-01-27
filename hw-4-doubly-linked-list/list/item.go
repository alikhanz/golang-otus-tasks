package list

// Item элемент списка
type Item struct {
	next *Item
	prev *Item
	value interface{}
}

// NewItem конструктор элемента
func NewItem(v interface{}) Item {
	return Item{value: v}
}

// Next ссылка на следующий элемент
func (i Item) Next() *Item {
	return i.next
}

// Prev ссылка на предыдущий элемент
func (i Item) Prev() *Item {
	return i.prev
}

// Value значение элемента
func (i Item) Value() interface{} {
	return i.value
}