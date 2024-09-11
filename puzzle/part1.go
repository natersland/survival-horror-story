package puzzle

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RiddlePuzzlePart1 interface {
	EscapeTheLockedDoor(combo string) int
	MirrorCipher(symbol string) int
	OpenTheMetallicDoor(code string) int
	UnlockTheVault(code string) int
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

func (r *riddlePuzzleStruct) OpenTheMetallicDoor(code string) int {
	// 	The door requires solving a complex numeric code, and only by deciphering the symbols can you open it.
	// **🧩 Difficulty:** ⭐⭐⭐
	// **📝 Input:** A string of numbers.
	// **📝 Output:** Return the sum of all the digits in the string.

	sum := 0
	for _, char := range code {
		digit, _ := strconv.Atoi(string(char))
		sum += digit
	}

	return sum
}

func (r *riddlePuzzleStruct) UnlockTheVault(code string) int {
	// **Unlock the Vault**
	// The vault door has a code lock with a sequence of numbers, but the code must be deciphered based on a pattern.
	// **🧩 Difficulty:** ⭐⭐
	// **📝 Input:** A sequence of numbers.
	// **📝 Output:** Return the product of all the digits in the sequence.
	re := regexp.MustCompile(`\D`) // \D matches any non-digit character
	formattedCode := re.ReplaceAllString(code, "")

	result := 1

	for _, ch := range formattedCode {
		digit, _ := strconv.Atoi(string(ch))
		result *= digit
	}

	return result
}
