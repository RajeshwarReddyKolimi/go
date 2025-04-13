package main

import "fmt"

type Q[T any] struct {
	elements []T
}

func New[T any]() Q[T] {
	return Q[T]{}
}

func (q *Q[T]) Enqueue(element T) {
	q.elements = append(q.elements, element)
}

func (q *Q[T]) Dequeue() (T, error) {
	var empty T
	if len(q.elements) == 0 {
		return empty, fmt.Errorf("Queue is empty")
	}
	removedElement := q.elements[0]
	q.elements = q.elements[1:]
	return removedElement, nil
}

func (q *Q[T]) Peek() (T, error) {
	var empty T
	if len(q.elements) == 0 {
		return empty, fmt.Errorf("Queue is empty")
	}
	removedElement := q.elements[0]
	return removedElement, nil
}

func main() {
	q1 := New[int]()
	q1.Enqueue(1)
	q1.Enqueue(2)
	q1.Enqueue(3)
	q1.Enqueue(4)
	removedInt, er := q1.Dequeue()
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("Removeded element:", removedInt)
	}
	peekInt, er := q1.Peek()
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("Peeked element:", peekInt)
	}

	q2 := New[string]()
	q2.Enqueue("abc")
	q2.Enqueue("def")
	q2.Enqueue("jkl")
	removedString, er := q2.Dequeue()
	q2.Enqueue("mno")

	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("Removeded element:", removedString)
	}
	peekString, er := q2.Peek()
	if er != nil {
		fmt.Println(er)
	} else {
		fmt.Println("Peeked element:", peekString)
	}

}
