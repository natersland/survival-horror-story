package puzzle

import (
	"fmt"
	"strings"
)

type RiddlePuzzlePart1 interface {
	EscapeTheLockedDoor(combo string) int
	MirrorCipher(symbol string) int
}

type riddlePuzzleStruct struct {
}

func NewRiddlePuzzlePart1() RiddlePuzzlePart1 {
	return &riddlePuzzleStruct{}
}

func (r *riddlePuzzleStruct) EscapeTheLockedDoor(combo string) int {
	// 	**Escape the Locked Door!**
	// You’re now inside the house, but the front door slams shut behind you. To escape, you must solve this riddle:
	// **🧩 Difficulty:** ⭐
	// **📝 Input:** A string representing the combination.
	// **📝 Output:** Return the number of characters in the string.
	noSpacesCombo := strings.ReplaceAll(combo, " ", "")
	countingCharLength := len(noSpacesCombo)

	return countingCharLength

}

func (r *riddlePuzzleStruct) MirrorCipher(symbol string) int {
	// **Mirror Cipher**
	// The mirror begins to glow brighter,
	// showing you strange symbols. You must decode them to unlock the secrets of the mansion.
	// **🧩 Difficulty:** ⭐⭐
	// **📝 Input:** A string of symbols (characters).
	// **📝 Output:** Return the number of unique symbols.
	noSpaceSymbol := strings.ReplaceAll(symbol, " ", "")
	symbolCount := make(map[rune]int)
	println("before loop: now symbol count is: ", len(symbolCount))

	for _, char := range noSpaceSymbol {
		symbolCount[char]++
		fmt.Printf("now symbol count for %c is: %d\n", char, len(symbolCount))
	}

	uniqueSymbol := len(symbolCount)

	return uniqueSymbol

}
