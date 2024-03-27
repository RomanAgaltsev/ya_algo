package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	palindromeTests := []struct{
		testNumber string
		line string
		isPalindrome bool
	}{
		{"1", "A man, a plan, a canal: Panama", true},
		{"2", "zo", false},
		{"3", "ThiS_String-is-@-PALIND0m3```~3m0DNILAP-()-si*!gnirts>>>siht", true},
		{"4", "-Luke, I'm your Father! -N00Oo! -oo00n -rehTAFruoymiekul", true},
	}
	
	for _, tt := range palindromeTests {
		t.Run(tt.testNumber, func(t *testing.T) {
			got := isPalindrome(tt.line)
			if got != tt.isPalindrome {
				t.Errorf("%q got %v want %v", tt.line, got, tt.isPalindrome)
			}
		})
	}
}