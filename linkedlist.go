package forms

type node struct {
	key  string
	next *node
}

func (n *node) equal(other *node) bool {
	return n.key == other.key
}

type linkedList struct {
	head *node
	tail *node
}

func newLinkedList() *linkedList {
	return &linkedList{}
}

func (l *linkedList) append(v string) {
	var node = &node{key: v}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
}

func (l *linkedList) remove(v string) bool {
	if l.head == nil {
		return false
	}
	if l.head.key == v {
		l.head = l.head.next
		if l.head == nil {
			l.tail = nil
		}
		return true
	}
	var n = l.head
	for n.next != nil && n.next.key != v {
		n = n.next
	}
	if n.next == l.tail {
		l.tail = n
	}
	if n.next != nil {
		n.next = n.next.next
		return true
	}
	return false
}
