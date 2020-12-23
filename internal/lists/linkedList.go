package lists

import (
	"fmt"
)

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

func NewNode(val int) *Node {
	n := Node{Value: val, Next: nil, Prev: nil}
	return &n
}

type LinkedList struct {
	table map[int]*Node
	Start *Node
}

func NewLinkedList() *LinkedList {
	M := map[int]*Node{}
	l := LinkedList{Start: nil, table: M}
	return &l
}
func (l *LinkedList) InsertEnd(val int) {
	newNode := NewNode(val)
	l.table[val] = newNode
	if l.Start == nil {
		newNode.Next = newNode
		newNode.Prev = newNode
		l.Start = newNode

	} else {
		last := l.Start.Prev
		newNode.Next = l.Start
		l.Start.Prev = newNode
		newNode.Prev = last
		last.Next = newNode
	}
}

func (l *LinkedList) InsertAfter(node *Node, val1, val2, val3 int) {
	n1 := NewNode(val1)
	n2 := NewNode(val2)
	n3 := NewNode(val3)
	l.table[val1] = n1
	l.table[val2] = n2
	l.table[val3] = n3

	tmp := node.Next
	node.Next = n1

	n1.Next = n2
	n1.Prev = node

	n2.Prev = n1
	n2.Next = n3

	n3.Prev = n2
	n3.Next = tmp

	l.table[tmp.Value] = tmp

	tmp.Prev = n3

}

func (l *LinkedList) DeleteAfter(node *Node, amount int) {
	tmp := node.Next
	l.table[tmp.Value] = nil

	if tmp == l.Start {
		l.Start = node
	}
	tmp.Next.Prev = node
	node.Next = tmp.Next

	if amount > 0 {
		l.DeleteAfter(node, amount-1)
	}
	//fmt.Println("START VAL ",l.Start.Value)
	// l.AdjustMinMax()
}

func (l *LinkedList) PrettyPrint() {
	temp := l.Start

	for temp.Next != l.Start {
		fmt.Printf("%d ", temp.Value)
		temp = temp.Next
	}
	fmt.Printf("%d\n", temp.Value)
}

func (l *LinkedList) ToArray() []int {
	var arr []int
	temp := l.Start

	for temp.Next != l.Start {
		arr = append(arr, temp.Value)
		temp = temp.Next
	}
	arr = append(arr, temp.Value)
	return arr

}

func (l *LinkedList) Find(target int) *Node {

	if n, ok := l.table[target]; ok && n != nil {
		return n
	}

	tmp := l.Start
	tmp2 := l.Start.Next

	for tmp != tmp2 {
		if tmp.Value == target {
			l.table[target] = tmp
			return tmp
		}
		if tmp2.Value == target {
			l.table[target] = tmp2
			return tmp2
		}
		tmp = tmp.Prev
		tmp2 = tmp2.Next
	}

	if tmp.Value == target {
		l.table[target] = tmp
		return tmp

	}

	if tmp2.Value == target {
		l.table[target] = tmp2
		return tmp2

	}

	return nil
}
