package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Next:  l.front,
	}
	if l.len == 0 {
		l.back = item
	} else {
		l.front.Prev = item
	}

	l.front = item
	l.len++

	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	item := &ListItem{
		Value: v,
		Prev:  l.back,
	}
	if l.len == 0 {
		l.front = item
	} else {
		l.back.Next = item
	}

	l.back = item
	l.len++

	return item
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil { // i не первый элемент
		i.Prev.Next = i.Next
	} else { // i первый элемент
		l.front = i.Next
	}

	if i.Next != nil { // i не последний элемент
		i.Next.Prev = i.Prev
	} else { // i последний элемент
		l.back = i.Prev
	}

	i.Next = nil
	i.Prev = nil
	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}
	l.Remove(i)

	i.Next = l.front
	i.Prev = nil

	if l.front != nil {
		l.front.Prev = i
	} else {
		l.back = i
	}

	l.front = i
	l.len++
}
