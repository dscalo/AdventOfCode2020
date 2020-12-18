package puzzles

import "fmt"

type Token struct {
	Type  string
	Value int64
}

type Tokens = []Token

func PrettyPrintTokens(ts Tokens) {
	for _, t := range ts {
		switch t.Type {
		case "A":
			fmt.Printf(" + ")
		case "M":
			fmt.Printf(" * ")
		case "N":
			fmt.Printf(" %d ", t.Value)
		case "LP":
			fmt.Printf(" (")
		case "RP":
			fmt.Printf(") ")
		}
	}
	fmt.Println("")
}

func NewToken(s string, v int64) *Token {
	t := Token{Type: s, Value: v}
	return &t
}

type node struct {
	token Token
	next  *node
}

type Stack struct {
	head *node
	size int
}

func (s *Stack) Peek() *Token {
	if s.size == 0 {
		return nil
	}
	t := s.head.token
	return &t
}

func (s *Stack) Push(token *Token) {
	s.size++
	n := node{token: *token, next: s.head}
	s.head = &n

}

func (s *Stack) Pop() *Token {
	if s.size == 0 {
		return nil
	}

	n := s.head
	t := n.token
	s.head = n.next
	s.size--

	return &t
}

func (s *Stack) HasNext() bool {
	return s.size > 0
}

func (s *Stack) PrettyPrint() {
	n := s.head

	for n != nil {
		switch n.token.Type {
		case "A":
			fmt.Printf(" + ")
		case "M":
			fmt.Printf(" * ")
		case "N":
			fmt.Printf(" %d ", n.token.Value)
		}
		n = n.next
	}
	fmt.Println("")
}

func NewStack() *Stack {
	s := Stack{head: nil, size: 0}
	return &s
}

type Queue struct {
	front *node
	rear  *node
}

func (q *Queue) Peek() *Token {
	if q.front == nil {
		return nil
	}
	t := q.front.token
	return &t
}

func (q *Queue) Enqueue(token *Token) {
	n := node{token: *token}
	if q.rear == nil {
		q.front = &n
		q.rear = &n
	}

	q.rear.next = &n
	q.rear = &n

}

func (q *Queue) Dequeue() *Token {
	if q.front == nil {
		return nil
	}

	n := q.front
	q.front = n.next

	if q.front == nil {
		q.rear = nil
	}

	return &n.token
}

func (q *Queue) Reverse() {
	stack := NewStack()

	for q.HasNext() {
		stack.Push(q.Dequeue())
	}

	for stack.HasNext() {
		q.Enqueue(stack.Pop())
	}
}

func (q Queue) HasNext() bool {
	return q.front != nil
}

func (q *Queue) PrettyPrint() {
	n := q.front

	for n != nil {
		switch n.token.Type {
		case "A":
			fmt.Printf(" + ")
		case "M":
			fmt.Printf(" * ")
		case "N":
			fmt.Printf(" %d ", n.token.Value)
		}
		n = n.next
	}
	fmt.Println("")
}

func NewQueue() *Queue {
	q := Queue{front: nil, rear: nil}
	return &q
}
