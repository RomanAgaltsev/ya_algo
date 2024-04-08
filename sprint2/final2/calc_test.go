package main

import "testing"

func TestCalc(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		calc := newCalculator()
		calc.exp = "2 1 + 3 *"
		got := calc.calculate()
		if got != 9 {
			t.Errorf("%q got %v want %v", calc.exp, got, 9)
		}
	})

	t.Run("2", func(t *testing.T) {
		calc := newCalculator()
		calc.exp = "7 2 + 4 * 2 +"
		got := calc.calculate()
		if got != 38 {
			t.Errorf("%q got %v want %v", calc.exp, got, 38)
		}
	})

	t.Run("23", func(t *testing.T) {
		calc := newCalculator()
		calc.exp = "2 5 - 4 /"
		got := calc.calculate()
		if got != -1 {
			t.Errorf("%q got %v want %v", calc.exp, got, -1)
		}
	})

	t.Run("28", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "-4 3 * 3 -5 - / -7 -1 - -8 -10 - + *"
		got := calc.calculate()
		if got != 8 {
			t.Errorf("%q got %v want %v", calc.exp, got, 8)
		}
	})

	t.Run("31", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "0 10 * -8 10 / + -9 4 / -10 5 * - * 1 0 - 6 -3 * - 7 3 / 10 -6 - - - /"
		got := calc.calculate()
		if got != -2 {
			t.Errorf("%q got %v want %v", calc.exp, got, -2)
		}
	})

	t.Run("33", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "-3 -7 - 1 -10 * * -5 9 / 6 1 * * + 9 2 / 3 3 / / 10 0 - 4 -10 - / * -"
		got := calc.calculate()
		if got != -46 {
			t.Errorf("%q got %v want %v", calc.exp, got, -46)
		}
	})

	t.Run("34", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "7 -9 + 9 0 - / 4 7 + -2 1 / + - 5 2 / 6 -8 * + 6 -3 - 4 8 + * / - -4 1 * 0 8 + / -9 -8 - 6 -4 * - / 0 8 - -9 3 + - -9 -4 * -6 6 * - / * +"
		got := calc.calculate()
		if got != -8 {
			t.Errorf("%q got %v want %v", calc.exp, got, -8)
		}
	})

	t.Run("36", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "-7 -7 + 9 8 / / -2 1 - 7 3 / / - -5 -9 - 7 -7 - / 5 4 + -8 1 * + + - 4 -4 * 6 5 - / -7 -7 + 7 2 - + * 10 10 + -7 1 / + 10 -1 - -4 0 * - / - / 3 8 + -2 6 + / 2 7 * -7 -9 - - / -7 2 / -6 -8 * / -7 2 / -6 -9 * * + * 8 1 / 0 2 * * -5 6 / -4 1 / * / 4 -10 * -5 2 * - 2 2 / 3 4 + / - - + +"
		got := calc.calculate()
		if got != 29 {
			t.Errorf("%q got %v want %v", calc.exp, got, 29)
		}
	})

	t.Run("37", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "1 8 / 3 -9 - / -1 7 + -10 -8 * + / -3 -9 + -7 0 + * 2 -8 * 2 7 / - - + 6 5 * 0 -8 * * -7 6 + -10 10 + + - -3 10 / 4 8 * * 7 -1 + -6 5 / - / - / 5 -7 + 3 4 / + 0 6 + 3 4 - - / -7 5 * 1 6 / * 2 10 / 1 10 + - + - -2 7 / -4 -7 * + 1 5 / 5 7 / * + 2 1 / -4 7 / + 6 4 * 4 9 / - + / + *"
		got := calc.calculate()
		if got != 220 {
			t.Errorf("%q got %v want %v", calc.exp, got, 220)
		}
	})

	t.Run("90", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "7 2 3 * -"
		got := calc.calculate()
		if got != 1 {
			t.Errorf("%q got %v want %v", calc.exp, got, 1)
		}
	})

	t.Run("91", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "2 3 * 4 5 * +"
		got := calc.calculate()
		if got != 26 {
			t.Errorf("%q got %v want %v", calc.exp, got, 26)
		}
	})

	t.Run("92", func(t *testing.T) {
		calc := newCalculator()

		calc.exp = "8 2 5 * + 1 3 2 * + 4 - /"
		got := calc.calculate()
		if got != 6 {
			t.Errorf("%q got %v want %v", calc.exp, got, 6)
		}
	})
}