package main

import "testing"

func TestHashTable(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		commands := []struct{
			command сommand
			result string
		}{
			{сommand{"get", 1, 0}, "None"},
			{сommand{"put", 1, 10}, ""},
			{сommand{"put", 2, 4}, ""},
			{сommand{"get", 1, 0}, "10"},
			{сommand{"get", 2, 0}, "4"},
			{сommand{"delete", 2, 0}, "4"},
			{сommand{"get", 2, 0}, "None"},
			{сommand{"put", 1, 5}, ""},
			{сommand{"get", 1, 0}, "5"},
			{сommand{"delete", 2, 0}, "None"},
		}
		hashTable := NewHashTable()
		for _, command := range commands {
			got := executeCommand(hashTable, command.command)
			if got != command.result {
				t.Errorf("%#v got %v want %v", command, got, command.result)
			}
		}
	})
	
	t.Run("2", func(t *testing.T) {
		commands := []struct{
			command сommand
			result string
		}{
			{сommand{"get", 9, 0}, "None"},
			{сommand{"delete", 9, 0}, "None"},
			{сommand{"put", 9, 1}, ""},
			{сommand{"get", 9, 0}, "1"},
			{сommand{"put", 9, 2}, ""},
			{сommand{"get", 9, 0}, "2"},
			{сommand{"put", 9, 3}, ""},
			{сommand{"get", 9, 0}, "3"},
		}
		hashTable := NewHashTable()
		for _, command := range commands {
			got := executeCommand(hashTable, command.command)
			if got != command.result {
				t.Errorf("%#v got %v want %v", command, got, command.result)
			}
		}
	})
}