package leap

import (
	"testing"
)

func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}


var testCases = []struct {
	year        int
	expected    bool
	description string
}{
	{1996, true, "vanilla leap year"},
	{1997, false, "normal year"},
	{1900, false, "century"},
	{2400, true, "exceptional century"},
}

func TestLeapYears(t *testing.T) {
	for _, test := range testCases {
		observed := IsLeapYear(test.year)
		if observed != test.expected {
			t.Fatalf("IsLeapYear(%d) = %t, want %t (%s)",
				test.year, observed, test.expected, test.description)
		}
	}
}

func BenchmarkLeapYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			IsLeapYear(test.year)
		}
	}
}
