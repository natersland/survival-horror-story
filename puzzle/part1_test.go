package puzzle

import (
	"testing"
)

func TestEscapeTheLockedDoor(t *testing.T) {
	// arrange
	riddle := NewRiddlePuzzlePart1()
	testCase := []struct {
		name   string
		combo  string
		expect int
	}{
		{
			name:   "Combo 1 - normal",
			combo:  "jashdkashdccc",
			expect: 13,
		},
		{
			name:   "Combo 2 - with space",
			combo:  "cat jaaaa",
			expect: 8,
		},
		{
			name:   "Combo 3 - with space trim case",
			combo:  " dog eieiei ",
			expect: 9,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			// act
			got := riddle.EscapeTheLockedDoor(tc.combo)

			//assert
			if got != tc.expect {
				t.Errorf("🚪:EscapeTheLockedDoor(%q) = %d, want %d", tc.combo, got, tc.expect)
			}
		})
	}
}

func TestMirrorCipher(t *testing.T) {
	// arrange
	riddle := NewRiddlePuzzlePart1()
	testCases := []struct {
		name   string
		symbol string
		expect int
	}{
		{
			name:   "Symbol 1 - normal",
			symbol: "aabcc",
			expect: 3,
		},
		{
			name:   "Symbol 1 - with space",
			symbol: "aab cc",
			expect: 3,
		},
		{
			name:   "Symbol 1 - with trim",
			symbol: " aab cc ",
			expect: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// act
			got := riddle.MirrorCipher(tc.symbol)

			// assert
			if got != tc.expect {
				t.Errorf("🪞:MirrorCipher(%q) = %d, want %d", tc.symbol, got, tc.expect)
			}
		})
	}
}
