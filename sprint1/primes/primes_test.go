package primes

import "testing"

var numsMap = map[int]bool {
	0: false,
	1: false,
	2: true,
	4: false,
	5: true,
	6: false,
	7: true,
	8: false,
	9: false,
	10: false,
	11: true,
	12: false,
	13: true,
	14: false,
	15: false,
	16: false,
	17: true,
	18: false,
	19: true,
	20: false,
	}

func TestIsPrime(t *testing.T) {
	for num, want := range numsMap {
		expected := IsPrime(num)
		if expected != want {
			t.Errorf("Common: num %v expected %v bot got %v", num, want, expected)
		}
	}
}

func TestIsPrimeFast(t *testing.T) {
	for num, want := range numsMap {
		expected := IsPrimeFast(num)
		if expected != want {
			t.Errorf("Fast: num %v expected %v bot got %v", num, want, expected)
		}
	}
}

func TestIsPrimeTricky(t *testing.T) {
	for num, want := range numsMap {
		expected := IsPrimeTricky(num)
		if expected != want {
			t.Errorf("Tricky: num %v expected %v bot got %v", num, want, expected)
		}
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for num, want := range numsMap {
		expected := IsPrime(num)
		if expected != want {
			b.Errorf("expected %v bot got %v", want, expected)
		}
	}
}

func BenchmarkIsPrimeFast(b *testing.B) {
	for num, want := range numsMap {
		expected := IsPrime(num)
		if expected != want {
			b.Errorf("expected %v bot got %v", want, expected)
		}
	}
}

func BenchmarkIsPrimeTricky(b *testing.B) {
	for num, want := range numsMap {
		expected := IsPrimeTricky(num)
		if expected != want {
			b.Errorf("expected %v bot got %v", want, expected)
		}
	}
}