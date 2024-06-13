package main

import (
	"reflect"
	"testing"
)

var assertCorrectMsg = func(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v want: %v", got, want)
	}
}

var assertCorrectData = func(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

var assertErrorMsg = func(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("expected an error but did not get one: %v", got)
	}
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestInsert(t *testing.T) {

	t.Run("insert 1 node at beginning in empty linked list", func(t *testing.T) {
		ll := LinkedList{}
		ll.insertBeginning(2)

		got := ll.Display()
		want := []int{2}

		assertCorrectMsg(t, got, want)
	})
	t.Run("insert 1 node at beginning", func(t *testing.T) {
		ll := LinkedList{}
		ll.insertBeginning(2, 4)

		got := ll.Display()
		want := []int{4, 2}

		assertCorrectMsg(t, got, want)
	})

	t.Run("insert 3 nodes at beginning", func(t *testing.T) {
		ll := LinkedList{}
		ll.insertBeginning(2, 3, 4)

		got := ll.Display()
		want := []int{4, 3, 2}

		assertCorrectMsg(t, got, want)
	})

	t.Run("insert 2 nodes at end", func(t *testing.T) {
		ll := LinkedList{}
		ll.insertEnd(5, 3)

		got := ll.Display()
		want := []int{5, 3}

		assertCorrectMsg(t, got, want)
	})
	t.Run("insert 1 node at end in empty linked list", func(t *testing.T) {
		ll := LinkedList{}
		ll.insertEnd(10)

		got := ll.Display()
		want := []int{10}

		assertCorrectMsg(t, got, want)
	})

}

func TestLength(t *testing.T) {
	ll := LinkedList{}
	ll.insertBeginning(1, 2, 3, 4, 5)

	got := ll.getLength()
	want := 5

	if got != want {
		t.Errorf("got: %v want: %v", got, want)
	}
}

func TestRemoveBeginning(t *testing.T) {

	ll := LinkedList{}
	t.Run("remove from empty linked list", func(t *testing.T) {
		err := ll.removeBeginning()
		assertErrorMsg(t, err, ErrEmptyLinkedList)
	})

	t.Run("remove 1 from beginning", func(t *testing.T) {
		ll.insertEnd(1, 2, 3)
		ll.removeBeginning()

		got := ll.Display()
		want := []int{2, 3}

		assertCorrectMsg(t, got, want)
	})

}

func TestAccessByIndex(t *testing.T) {
	ll := LinkedList{}
	ll.insertEnd(1, 2, 3, 4, 5)

	t.Run("access index 0", func(t *testing.T) {
		node, _ := ll.accessByIndex(0)
		assertCorrectData(t, node.data, 1)
	})

	t.Run("access last index", func(t *testing.T) {
		last_idx := ll.getLength() - 1
		node, _ := ll.accessByIndex(last_idx)
		assertCorrectData(t, node.data, 5)
	})

	t.Run("test out of index", func(t *testing.T) {
		_, err := ll.accessByIndex(9)
		assertErrorMsg(t, err, ErrOutOfRange)
	})
}
