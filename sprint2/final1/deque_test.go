package main

import "testing"

func TestDeque(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		deque := NewDeque(4)
		
		err := deque.PushFront(861)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		
		err = deque.PushFront(-819)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		
		value, err := deque.PopBack()
		if value != 861 {
			t.Errorf("PopBack: got %v want %v", value, 861)
		}
		if err != nil {
			t.Errorf("PopBack: got %v want %v", err, nil)
		}
		
		value, err = deque.PopBack()
		if value != -819 {
			t.Errorf("PopBack: got %v want %v", value, -819)
		}
		if err != nil {
			t.Errorf("PopBack: got %v want %v", err, nil)
		}
	})
	
	t.Run("2", func(t *testing.T) {
		deque := NewDeque(10)
		
		err := deque.PushFront(-855)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		
		err = deque.PushFront(0)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		
		value, err := deque.PopBack()
		if value != -855 {
			t.Errorf("PopBack: got %v want %v", value, -855)
		}
		if err != nil {
			t.Errorf("PopBack: got %v want %v", err, nil)
		}
		
		value, err = deque.PopBack()
		if value != 0 {
			t.Errorf("PopBack: got %v want %v", value, 0)
		}
		if err != nil {
			t.Errorf("PopBack: got %v want %v", err, nil)
		}
		
		err = deque.PushBack(844)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		
		value, err = deque.PopBack()
		if value != 844 {
			t.Errorf("PopBack: got %v want %v", value, 844)
		}
		if err != nil {
			t.Errorf("PopBack: got %v want %v", err, nil)
		}
		
		err = deque.PushBack(823)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
	})
	
	t.Run("3", func(t *testing.T) {
		deque := NewDeque(6)
		
		err := deque.PushFront(-201)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		
		err = deque.PushBack(959)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		
		err = deque.PushBack(102)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		
		err = deque.PushFront(20)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		
		value, err := deque.PopFront()
		if value != 20 {
			t.Errorf("PopFront: got %v want %v", value, 20)
		}
		if err != nil {
			t.Errorf("PopFront: got %v want %v", err, nil)
		}
		
		value, err = deque.PopBack()
		if value != 102 {
			t.Errorf("PopBack: got %v want %v", value, 102)
		}
		if err != nil {
			t.Errorf("PopBack: got %v want %v", err, nil)
		}
	})
	
	t.Run("12", func(t *testing.T) {
		deque := NewDeque(8)
		
		err := deque.PushFront(842)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		value, err := deque.PopBack()
		if value != 842 {
			t.Errorf("PopBack: got %v want %v", value, 842)
		}
		if err != nil {
			t.Errorf("PopBack: got %v want %v", err, nil)
		}
		err = deque.PushFront(576)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		err = deque.PushFront(-853)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		value, err = deque.PopBack()
		if value != 576 {
			t.Errorf("PopBack: got %v want %v", value, 576)
		}
		err = deque.PushFront(123)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		err = deque.PushFront(-236)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		value, err = deque.PopFront()
		if value != -236 {
			t.Errorf("PopFront: got %v want %v", value, -236)
		}
		err = deque.PushBack(840)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		value, err = deque.PopFront()
		if value != 123 {
			t.Errorf("PopFront: got %v want %v", value, 123)
		}
		err = deque.PushBack(740)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		err = deque.PushBack(347)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		value, err = deque.PopFront()
		if value != -853 {
			t.Errorf("PopFront: got %v want %v", value, -853)
		}
		err = deque.PushFront(-767)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		err = deque.PushFront(-711)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		err = deque.PushBack(-7)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		value, err = deque.PopFront()
		if value != -711 {
			t.Errorf("PopFront: got %v want %v", value, -711)
		}
		value, err = deque.PopBack()
		if value != -7 {
			t.Errorf("PopBack: got %v want %v", value, -7)
		}
		value, err = deque.PopBack()
		if value != 347 {
			t.Errorf("PopBack: got %v want %v", value, 347)
		}
		value, err = deque.PopBack()
		if value != 740 {
			t.Errorf("PopBack: got %v want %v", value, 740)
		}
		value, err = deque.PopBack()
		if value != 840 {
			t.Errorf("PopBack: got %v want %v", value, 840)
		}
		value, err = deque.PopFront()
		if value != -767 {
			t.Errorf("PopFront: got %v want %v", value, -767)
		}
		err = deque.PushBack(-215)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		err = deque.PushFront(540)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		value, err = deque.PopFront()
		if value != 540 {
			t.Errorf("PopFront: got %v want %v", value, 540)
		}
		err = deque.PushFront(-293)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		value, err = deque.PopBack()
		if value != -215 {
			t.Errorf("PopBack: got %v want %v", value, -215)
		}
		value, err = deque.PopBack()
		if value != -293 {
			t.Errorf("PopBack: got %v want %v", value, -293)
		}
		_, err = deque.PopFront()
		if err == nil {
			t.Errorf("PopFront: got %v want %v", err, ErrDequeIsEmpty)
		}
		_, err = deque.PopBack()
		if err == nil {
			t.Errorf("PopBack: got %v want %v", err, ErrDequeIsEmpty)
		}
		err = deque.PushBack(873)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		err = deque.PushFront(47)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		err = deque.PushBack(-238)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		err = deque.PushFront(-575)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		value, err = deque.PopFront()
		if value != -575 {
			t.Errorf("PopFront: got %v want %v", value, -575)
		}
		err = deque.PushFront(-916)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		err = deque.PushFront(292)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		err = deque.PushBack(302)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		err = deque.PushFront(456)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		err = deque.PushBack(92)
		if err != nil {
			t.Errorf("PushBack: got %v want %v", err, nil)
		}
		err = deque.PushBack(-422)
		if err == nil {
			t.Errorf("PushBack: got %v want %v", err, ErrDequeIsFull)
		}
		err = deque.PushBack(890)
		if err == nil {
			t.Errorf("PushBack: got %v want %v", err, ErrDequeIsFull)
		}
		err = deque.PushBack(-100)
		if err == nil {
			t.Errorf("PushBack: got %v want %v", err, ErrDequeIsFull)
		}
		value, err = deque.PopBack()
		if value != 92 {
			t.Errorf("PopBack: got %v want %v", value, 92)
		}
		err = deque.PushFront(-356)
		if err != nil {
			t.Errorf("PushFront: got %v want %v", err, nil)
		}
		value, err = deque.PopFront()
		if value != -356 {
			t.Errorf("PopFront: got %v want %v", value, -356)
		}
	})
}