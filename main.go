package main

import (
	"errors"
)

var (
	ErrOutOfRange      = errors.New("index out of range")
	ErrEmptyLinkedList = errors.New("empty linked list")
)

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Display() []int {
	curr := l.head
	display := []int{}

	for curr != nil {
		display = append(display, curr.data)
		curr = curr.next
	}

	return display

}

func (l *LinkedList) insertBeginning(d ...int) {

	for _, data := range d {
		node := &Node{data: data, next: l.head}
		l.head = node
	}

}

func (l *LinkedList) insertEnd(d ...int) {

	for _, data := range d {
		node := &Node{data: data}
		curr := l.head

		if curr == nil {
			l.head = node
		} else {
			// go to the last node of linked list and add node
			for curr.next != nil {
				curr = curr.next
			}
			curr.next = node
		}
	}

}

func (l *LinkedList) getLength() int {
	length := 0
	curr := l.head

	for curr != nil {
		length += 1
		curr = curr.next
	}

	return length
}

func (l *LinkedList) removeBeginning() error {
	if l.head == nil {
		return ErrEmptyLinkedList
	}
	l.head = l.head.next
	return nil
}

func (l *LinkedList) accessByIndex(index int) (*Node, error) {

	currIndex := 0
	curr := l.head

	for curr != nil {
		if currIndex == index {
			return curr, nil
		} else {
			curr = curr.next
			currIndex += 1
		}
	}

	return nil, ErrOutOfRange
}
