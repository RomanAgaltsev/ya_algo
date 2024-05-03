package main

import "testing"

func TestHashTable(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		commands := []struct{
			command Command
			result string
		}{
			{Command{"get", 1, 0}, "None"},
			{Command{"put", 1, 10}, ""},
			{Command{"put", 2, 4}, ""},
			{Command{"get", 1, 0}, "10"},
			{Command{"get", 2, 0}, "4"},
			{Command{"delete", 2, 0}, "4"},
			{Command{"get", 2, 0}, "None"},
			{Command{"put", 1, 5}, ""},
			{Command{"get", 1, 0}, "5"},
			{Command{"delete", 2, 0}, "None"},
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
			command Command
			result string
		}{
			{Command{"get", 9, 0}, "None"},
			{Command{"delete", 9, 0}, "None"},
			{Command{"put", 9, 1}, ""},
			{Command{"get", 9, 0}, "1"},
			{Command{"put", 9, 2}, ""},
			{Command{"get", 9, 0}, "2"},
			{Command{"put", 9, 3}, ""},
			{Command{"get", 9, 0}, "3"},
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