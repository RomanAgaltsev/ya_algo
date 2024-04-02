package main

import "testing"

func TestStackMaxEffective(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		stack := NewStack()

		_, err := stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.Peek()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}

		stack.Push(4)
		stack.Push(-5)

		item, err := stack.Peek()
		if item != -5 {
			t.Errorf("got %v want %v", item, -5)
		}

		stack.Push(7)
		stack.Pop()
		stack.Pop()

		item, err = stack.GetMax()
		if item != 4 {
			t.Errorf("got %v want %v", item, 4)
		}
		item, err = stack.Peek()
		if item != 4 {
			t.Errorf("got %v want %v", item, 4)
		}

		stack.Pop()

		item, err = stack.GetMax()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
	})

	t.Run("2", func(t *testing.T) {
		stack := NewStack()

		_, err := stack.GetMax()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}

		stack.Push(-6)
		stack.Pop()

		_, err = stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.GetMax()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}

		stack.Push(2)

		item, err := stack.GetMax()
		if item != 2 {
			t.Errorf("got %v want %v", item, 2)
		}

		stack.Pop()

		stack.Push(-2)
		stack.Push(-6)
	})

	t.Run("3", func(t *testing.T) {
		stack := NewStack()

		_, err := stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.GetMax()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}

		stack.Push(-6)
		stack.Pop()

		_, err = stack.GetMax()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}
		_, err = stack.Pop()
		if err == nil {
			t.Errorf("got %v want %v", err, ErrStackIsEmpty)
		}

		stack.Push(-6)

		item, err := stack.GetMax()
		if item != -6 {
			t.Errorf("got %v want %v", item, -6)
		}
	})
	
	t.Run("3", func(t *testing.T) {
		stack := NewStack()
		
		stack.Push(1)
		stack.Push(3)
		stack.Push(1)
		stack.Push(3)
		stack.Push(3)
		stack.Pop()
		
		item, _ := stack.GetMax()
		if item != 3 {
			t.Errorf("got %v want %v", item, 3)
		}
		
		stack.Pop()
		
		item, _ = stack.GetMax()
		if item != 3 {
			t.Errorf("got %v want %v", item, 3)
		}
		
		stack.Pop()
		
		item, _ = stack.GetMax()
		if item != 3 {
			t.Errorf("got %v want %v", item, 3)
		}
		
		stack.Pop()
		
		item, _ = stack.GetMax()
		if item != 1 {
			t.Errorf("got %v want %v", item, 1)
		}
	})
}