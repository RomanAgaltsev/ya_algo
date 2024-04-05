package main

import "testing"

func TestDeque(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		commands := []struct{
			command Command
			result string
		}{
			{Command{"push_front", 861}, ""},
			{Command{"push_front", -819}, ""},
			{Command{"pop_back", 0}, "861"},
			{Command{"pop_back", 0}, "-819"},
		}
		deque := newDeque(4)
		for _, command := range commands {
			got := executeCommand(deque, command.command)
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
			{Command{"push_front", -855}, ""},
			{Command{"push_front", 0}, ""},
			{Command{"pop_back", 0}, "-855"},
			{Command{"pop_back", 0}, "0"},
			{Command{"push_back", 844}, ""},
			{Command{"pop_back", 0}, "844"},
			{Command{"push_back", 823}, ""},
		}
			deque := newDeque(10)
			for _, command := range commands {
				got := executeCommand(deque, command.command)
				if got != command.result {
					t.Errorf("%#v got %v want %v", command, got, command.result)
				}
			}
	})

	t.Run("3", func(t *testing.T) {
		commands := []struct{
			command Command
			result string
		}{
			{Command{"push_front", -201}, ""},
			{Command{"push_back", 959}, ""},
			{Command{"push_back", 102}, ""},
			{Command{"push_front", 20}, ""},
			{Command{"pop_front", 0}, "20"},
			{Command{"pop_back", 0}, "102"},
		}
			deque := newDeque(6)
			for _, command := range commands {
				got := executeCommand(deque, command.command)
				if got != command.result {
					t.Errorf("%#v got %v want %v", command, got, command.result)
				}
			}
	})

	t.Run("12", func(t *testing.T) {
		commands := []struct{
			command Command
			result string
		}{
			{Command{"push_front", 842}, ""},
			{Command{"pop_back", 0}, "842"},
			{Command{"push_front", 576}, ""},
			{Command{"push_front", -853}, ""},
			{Command{"pop_back", 0}, "576"},
			{Command{"push_front", 123}, ""},
			{Command{"push_front", -236}, ""},
			{Command{"pop_front", 0}, "-236"},
			{Command{"push_back", 840}, ""},
			{Command{"pop_front", 0}, "123"},
			{Command{"push_back", 740}, ""},
			{Command{"push_back", 347}, ""},
			{Command{"pop_front", 0}, "-853"},
			{Command{"push_front", -767}, ""},
			{Command{"push_front", -711}, ""},
			{Command{"push_back", -7}, ""},
			{Command{"pop_front", 0}, "-711"},
			{Command{"pop_back", 0}, "-7"},
			{Command{"pop_back", 0}, "347"},
			{Command{"pop_back", 0}, "740"},
			{Command{"pop_back", 0}, "840"},
			{Command{"pop_front", 0}, "-767"},
			{Command{"push_back", -215}, ""},
			{Command{"push_front", 540}, ""},
			{Command{"pop_front", 0}, "540"},
			{Command{"push_front", -293}, ""},
			{Command{"pop_back", 0}, "-215"},
			{Command{"pop_back", 0}, "-293"},
			{Command{"pop_front", 0}, "error"},
			{Command{"pop_back", 0}, "error"},
			{Command{"push_back", 873}, ""},
			{Command{"push_front", 47}, ""},
			{Command{"push_back", -238}, ""},
			{Command{"push_front", -575}, ""},
			{Command{"pop_front", 0}, "-575"},
			{Command{"push_front", -916}, ""},
			{Command{"push_front", 292}, ""},
			{Command{"push_back", 302}, ""},
			{Command{"push_front", 456}, ""},
			{Command{"push_back", 92}, ""},
			{Command{"push_back", -422}, "error"},
			{Command{"push_back", 890}, "error"},
			{Command{"push_back", -100}, "error"},
			{Command{"pop_back", 0}, "92"},
			{Command{"push_front", -356}, ""},
			{Command{"pop_front", 0}, "-356"},
			{Command{"push_front", 430}, ""},
			{Command{"push_front", 469}, "error"},
			{Command{"push_back", -56}, "error"},
			{Command{"push_front", 273}, "error"},
			{Command{"pop_back", 0}, "302"},
			{Command{"pop_back", 0}, "-238"},
			{Command{"push_front", -397}, ""},
			{Command{"push_back", 131}, ""},
			{Command{"pop_front", 0}, "-397"},
			{Command{"push_back", -4}, ""},
			{Command{"push_front", -265}, "error"},
			{Command{"push_back", 10}, "error"},
			{Command{"push_back", -531}, "error"},
			{Command{"pop_back", 0}, "-4"},
			{Command{"pop_back", 0}, "131"},
			{Command{"pop_back", 0}, "873"},
			{Command{"push_front", 766}, ""},
			{Command{"pop_front", 0}, "766"},
			{Command{"push_front", -520}, ""},
			{Command{"pop_back", 0}, "47"},
			{Command{"pop_back", 0}, "-916"},
			{Command{"pop_back", 0}, "292"},
			{Command{"pop_back", 0}, "456"},
			{Command{"pop_back", 0}, "430"},
			{Command{"push_front", -386}, ""},
			{Command{"push_back", -320}, ""},
			{Command{"push_back", 21}, ""},
			{Command{"pop_front", 0}, "-386"},
			{Command{"push_back", 495}, ""},
			{Command{"pop_front", 0}, "-520"},
			{Command{"push_front", -95}, ""},
			{Command{"pop_front", 0}, "-95"},
			{Command{"pop_back", 0}, "495"},
			{Command{"push_back", 908}, ""},
			{Command{"push_front", 115}, ""},
			{Command{"pop_front", 0}, "115"},
		}
			deque := newDeque(8)
			for _, command := range commands {
				got := executeCommand(deque, command.command)
				if got != command.result {
					t.Errorf("%v got %v want %v", command, got, command.result)
				}
			}
	})
}