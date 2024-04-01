package main

import "testing"



func TestGetMax(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		stackMax := NewStack()
		
		got, err := stackMax.GetMax()
		if err == nil {
			t.Errorf("got %v want %v", got, ErrStackIsEmpty)
		}
		
		stackMax.Push(7)
		stackMax.Pop()
		stackMax.Push(-2)
		stackMax.Push(-1)
		stackMax.Pop()
		
		got, _ = stackMax.GetMax()
		if got != -2 {
			t.Errorf("got %v want %v", got, -2)
		}
		
		got, _ = stackMax.GetMax()
		if got != -2 {
			t.Errorf("got %v want %v", got, -2)
		}
	})
	
	t.Run("2", func(t *testing.T) {
		stackMax := NewStack()
		
		got, err := stackMax.GetMax()
		if err == nil {
			t.Errorf("got %v want %v", got, ErrStackIsEmpty)
		}
		
		got, err = stackMax.Pop()
		if err == nil {
			t.Errorf("got %v want %v", got, ErrStackIsEmpty)
		}
		
		got, err = stackMax.Pop()
		if err == nil {
			t.Errorf("got %v want %v", got, ErrStackIsEmpty)
		}
		
		got, err = stackMax.Pop()
		if err == nil {
			t.Errorf("got %v want %v", got, ErrStackIsEmpty)
		}
				
		stackMax.Push(10)
		
		got, _ = stackMax.GetMax()
		if got != 10 {
			t.Errorf("got %v want %v", got, 10)
		}
		
		stackMax.Push(-9)
	})
}