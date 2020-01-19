package list

import "errors"

// List двусвязный список
type List struct {
	first *Item
	last *Item
	items map[*Item]*Item
}

// NewList конструктор списка
func NewList() List {
	return List{
		items: make(map[*Item]*Item),
	}
}

// Len возвращает размер списка
func (l *List) Len() int {
	return len(l.items)
}

// First возвращает первый элемент списка
func (l *List) First() *Item {
	return l.first
}

// Last возвращает последний элемент списка
func (l *List) Last() *Item {
	return l.last
}

// PushFront добавляет элемент в начало списка
func (l *List) PushFront(v interface{}) {
	item := NewItem(v)
	itemPtr := &item

	if l.first != nil {
		l.first.prev = itemPtr
		item.next = l.first
	}

	if l.last == nil {
		l.last = itemPtr
	}
	l.first = itemPtr
	l.items[itemPtr] = itemPtr
}

// PushBack добавляет элемент в конец списка
func (l *List) PushBack(v interface{}) {
	item := NewItem(v)
	itemPtr := &item

	if l.last != nil {
		l.last.next = itemPtr
		item.prev = l.last
	}

	if l.first == nil {
		l.first = itemPtr
	}

	l.last = itemPtr
	l.items[itemPtr] = itemPtr
}

// Remove удаляет элемент из списка
func (l *List) Remove(i *Item) error {
	item, ok := l.items[i]

	if !ok {
		return errors.New("элемент не существует в списке")
	}

	if l.First() == item {
		// Смещаем второй элемент вначало, убираем в нем ссылку на предыдущий элемент.
		secondItem := item.next
		secondItem.prev = nil
		l.first = secondItem
	} else if l.Last() == item {
		// Смещаем предпоследний элемент в конец, и убираем в нем ссылку на следующий элемент
		prevLastItem := item.prev
		prevLastItem.next = nil
		l.last = prevLastItem
	} else {
		// Удаляем элемент из середины просто перелинковав элементы.
		item.next.prev = item.prev
		item.prev.next = item.next
	}

	delete(l.items, i)
	return nil
}