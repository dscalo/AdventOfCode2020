package puzzles

import "testing"

func Test_empty_stack(t *testing.T) {
	stack := NewStack()

	if stack.HasNext() {
		t.Error("Stack hasNext should be false")
	}

	if stack.Peek() != nil {
		t.Error("Stack peek should be nil")
	}

	if stack.Pop() != nil {
		t.Error("Stack pop should be nil")
	}
}

func Test_peek_stack(t *testing.T) {
	stack := NewStack()

	t1 := NewToken("token_1", 1)
	t2 := NewToken("token_2", 2)

	stack.Push(t1)
	stack.Push(t2)

	if !stack.HasNext() {
		t.Error("Stack hasNext should be true")
	}

	tok := stack.Peek()

	if tok.Type != "token_2" {
		t.Error("Type should be token_2")
	}
}

func Test_pop_stack(t *testing.T) {
	stack := NewStack()

	t1 := NewToken("token_1", 1)
	t2 := NewToken("token_2", 2)

	stack.Push(t1)
	stack.Push(t2)

	tok := stack.Pop()

	if tok.Type != "token_2" {
		t.Error("Type should be token_2")
	}
	tok = stack.Pop()

	if tok.Type != "token_1" {
		t.Error("Type should be token_1")
	}

	if stack.HasNext() {
		t.Error("stack should be empty")
	}
}

func Test_empty_queue(t *testing.T) {
	queue := NewQueue()

	if queue.HasNext() {
		t.Error("queue hasNext should be false")
	}

	if queue.Peek() != nil {
		t.Error("queue peek should be nil")
	}

	if queue.Dequeue() != nil {
		t.Error("queue dequeue should be nil")
	}
}

func Test_peek_queue(t *testing.T) {
	queue := NewQueue()

	t1 := NewToken("token_1", 1)
	t2 := NewToken("token_2", 2)

	queue.Enqueue(t1)
	queue.Enqueue(t2)

	if !queue.HasNext() {
		t.Error("queue hasNext should be true")
	}

	tok := queue.Peek()

	if tok.Type != "token_1" {
		t.Error("Type should be token_1")
	}
}

func Test_dequeue_queue(t *testing.T) {
	queue := NewQueue()

	t1 := NewToken("token_1", 1)
	t2 := NewToken("token_2", 2)

	queue.Enqueue(t1)
	queue.Enqueue(t2)

	tok := queue.Dequeue()

	if tok.Type != "token_1" {
		t.Error("Type should be token_1")
	}
	tok = queue.Dequeue()

	if tok.Type != "token_2" {
		t.Error("Type should be token_2")
	}

	if queue.HasNext() {
		t.Error("queue should be empty")
	}
}
