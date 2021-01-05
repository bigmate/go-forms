package forms

import (
	"testing"
)

func mock(names []string) *linkedList {
	var ll = newLinkedList()
	for i := 0; i < len(names); i++ {
		ll.append(names[i])
	}
	return ll
}

func Test_linkedList_append(t *testing.T) {
	var names = []string{"A", "B", "C", "D", "E", "F"}
	var ll = mock(names)
	var node = ll.head
	var i int
	for node != nil {
		if node.key != names[i] {
			t.Errorf("Order Violation: %s != %s", node.key, names[i])
		}
		node = node.next
		i++
	}
	if ll.tail.key != names[len(names)-1] {
		t.Errorf("Wrong tail")
		t.FailNow()
	}
	for ll.tail != nil {
		ll.remove(ll.tail.key)
	}
	ll.append("head")
	if ll.head == nil {
		t.Error("Head is not set")
		t.FailNow()
	}
	if ll.tail == nil {
		t.Error("Tail is not set")
		t.FailNow()
	}
	if ll.head != ll.tail {
		t.Errorf("Head != Tail: %s, %s", ll.head.key, ll.tail.key)
		t.FailNow()
	}

	var ll2 = mock([]string{"TEST"})
	ll2.remove("TEST")
	if ll2.tail != nil && ll2.head != nil {
		t.Fatal("Remove failed")
	}
	ll2.append("A")
	if ll2.head != ll2.tail || ll2.tail == nil {
		t.Error("New node has not been set")
	}
}

func Test_linkedList_remove(t *testing.T) {
	var names = []string{"A", "B", "C", "D", "E", "F"}
	var ll = mock(names)
	var tail = *ll.tail
	ll.remove(tail.key)
	if tail.equal(ll.tail) {
		t.Errorf("Tail has not been removed: %s == %s", ll.tail.key, tail.key)
		t.FailNow()
	}
	var head = *ll.head
	ll.remove(head.key)
	if head.equal(ll.head) {
		t.Errorf("Head has not been removed: %s == %s", head.key, ll.head.key)
		t.FailNow()
	}
	var mid = "D"
	ll.remove(mid)
	var node = ll.head
	for node != nil {
		if node.key == mid {
			t.Errorf("Node has not been removed: %s", node.key)
			t.FailNow()
		}
		node = node.next
	}
	for ll.tail != nil {
		ll.remove(ll.tail.key)
	}
	if ll.head != nil {
		t.Errorf("Head has not been removed: %s", ll.head.key)
	}
}
