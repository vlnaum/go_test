package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int
	Front() *listItem
	Back() *listItem
	PushFront(v interface{}) *listItem
	PushBack(v interface{}) *listItem
	Remove(i *listItem)
	MoveToFront(i *listItem)
}

type listItem struct {
	Value      interface{}
	Prev, Next *listItem
}

type list struct {
	count      int
	head, tail *listItem
}

func (l list) Len() int {
	return l.count
}

func (l list) Front() *listItem {
	return l.head
}

func (l list) Back() *listItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *listItem {
	node := listItem{v, nil, nil}

	if l.head == nil {
		l.tail, l.head = &node, &node
	} else {
		node.Prev = l.head
		node.Prev.Next = &node
		l.head = &node
	}

	l.count++
	return &node
}

func (l *list) PushBack(v interface{}) *listItem {
	node := listItem{v, nil, nil}

	if l.head == nil {
		l.tail, l.head = &node, &node
	} else {
		node.Next = l.tail
		node.Next.Prev = &node
		l.tail = &node
	}

	l.count++
	return &node
}

func (l *list) Remove(node *listItem) {
	switch {
	case node == l.head && node == l.tail:
		l.head, l.tail = nil, nil
	case node == l.head:
		l.head = l.head.Prev
		l.head.Next = nil
	case node == l.tail:
		l.tail = l.tail.Next
		l.tail.Prev = nil
	default:
		node.Next.Prev = node.Prev
		node.Prev.Next = node.Next
	}

	l.count--
}

func (l *list) MoveToFront(node *listItem) {
	switch node {
	case l.head:
		return
	case l.tail:
		node.Next.Prev = nil
		l.tail = node.Next
	default:
		node.Next.Prev = node.Prev
		node.Prev.Next = node.Next
	}

	node.Prev = l.head
	node.Next = nil
	l.head.Next = node
	l.head = node
}

func NewList() List {
	return &list{}
}
