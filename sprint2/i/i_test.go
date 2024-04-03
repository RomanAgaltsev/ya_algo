package main

import "testing"

func TestMyQueueSized(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		queue := NewQueue(2)
		
		_, err := queue.Peek()
		if err == nil {
			t.Errorf("Peek: got %v want %v", err, ErrQueueIsEmpty)
		}
		
		queue.Push(5)
		queue.Push(2)
		
		item, err := queue.Peek()
		if item != 5 {
			t.Errorf("Peek: got %v want %v", item, 5)
		}
		if queue.size != 2 {
			t.Errorf("Size: got %v want %v", queue.size, 2)
		}
		if queue.size != 2 {
			t.Errorf("Size: got %v want %v", queue.size, 2)
		}
		
		err = queue.Push(1) 
		if err == nil {
			t.Errorf("Push: got %v want %v", err, ErrQueueIsFull)
		}
		if queue.size != 2 {
			t.Errorf("Size: got %v want %v", queue.size, 2)
		}
		
	})
	
	t.Run("2", func(t *testing.T) {
		queue := NewQueue(1)
		
		queue.Push(1)
		if queue.size != 1 {
			t.Errorf("Size: got %v want %v", queue.size, 1)
		}
		
		err := queue.Push(3) 
		if err == nil {
			t.Errorf("Push: got %v want %v", err, ErrQueueIsFull)
		}
		if queue.size != 1 {
			t.Errorf("Size: got %v want %v", queue.size, 1)
		}
		
		err = queue.Push(1) 
		if err == nil {
			t.Errorf("Push: got %v want %v", err, ErrQueueIsFull)
		}
		
		item, _ := queue.Pop()
		if item != 1 {
			t.Errorf("Pop: got %v want %v", item, 1)
		}
		
		err = queue.Push(1) 
		if err != nil {
			t.Errorf("Push: got %v want %v", err, nil)
		}
		
		item, _ = queue.Pop()
		if item != 1 {
			t.Errorf("Pop: got %v want %v", item, 1)
		}
		
		err = queue.Push(3) 
		if err != nil {
			t.Errorf("Push: got %v want %v", err, nil)
		}
		
		err = queue.Push(3) 
		if err == nil {
			t.Errorf("Push: got %v want %v", err, ErrQueueIsFull)
		}
	})
	
	
}