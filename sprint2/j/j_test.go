package main

import "testing"

func TestLinkedListQueue(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		queue := NewQueue()
		
		queue.Put(-34)
		queue.Put(-23)
		
		item, _ := queue.Get()
		if item != -34 {
			t.Errorf("Get: got %v want %v", item, -34)
		}
		if queue.size != 1 {
			t.Errorf("Size: got %v want %v", queue.size, 1)
		}
		
		item, _ = queue.Get()
		if item != -23 {
			t.Errorf("Get: got %v want %v", item, -23)
		}
		if queue.size != 0 {
			t.Errorf("Size: got %v want %v", queue.size, 0)
		}
		
		_, err := queue.Get()
		if err == nil {
			t.Errorf("Get: got %v want %v", err, ErrQueueIsEmpty)
		}
		_, err = queue.Get()
		if err == nil {
			t.Errorf("Get: got %v want %v", err, ErrQueueIsEmpty)
		}
		
		queue.Put(80)
		
		if queue.size != 1 {
			t.Errorf("Size: got %v want %v", queue.size, 1)
		}
	})
	
	t.Run("2", func(t *testing.T) {
		queue := NewQueue()
		
		queue.Put(-66)
		queue.Put(98)
		
		if queue.size != 2 {
			t.Errorf("Size: got %v want %v", queue.size, 2)
		}
		if queue.size != 2 {
			t.Errorf("Size: got %v want %v", queue.size, 2)
		}
		
		item, _ := queue.Get()
		if item != -66 {
			t.Errorf("Get: got %v want %v", item, -66)
		}
		
		item, _ = queue.Get()
		if item != 98 {
			t.Errorf("Get: got %v want %v", item, 98)
		}
	})
	
	t.Run("3", func(t *testing.T) {
		queue := NewQueue()
		
		_, err := queue.Get()
		if err == nil {
			t.Errorf("Get: got %v want %v", err, ErrQueueIsEmpty)
		}
		if queue.size != 0 {
			t.Errorf("Size: got %v want %v", queue.size, 0)
		}
		
		queue.Put(74)
		
		item, _ := queue.Get()
		if item != 74 {
			t.Errorf("Get: got %v want %v", item, 74)
		}
		if queue.size != 0 {
			t.Errorf("Size: got %v want %v", queue.size, 0)
		}
		
		queue.Put(90)
		
		if queue.size != 1 {
			t.Errorf("Size: got %v want %v", queue.size, 1)
		}
		if queue.size != 1 {
			t.Errorf("Size: got %v want %v", queue.size, 1)
		}
		if queue.size != 1 {
			t.Errorf("Size: got %v want %v", queue.size, 1)
		}
	})
}