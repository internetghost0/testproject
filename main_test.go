package main
import "testing"

func TestConvertRomanNumeralTo1(t *testing.T) {

	n := RomanToNumber("I")

	if n != 1 {
		t.Errorf("I should return 1, but %v", n)
	}
}
