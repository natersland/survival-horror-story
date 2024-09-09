package main

import (
	"github.com/natersland/survival-horror-story/puzzle"
)

func main() {
	riddleP1 := puzzle.NewRiddlePuzzlePart1()
	escapeTheLockedDoor := riddleP1.EscapeTheLockedDoor("4f2d6")
	mirrorCipher := riddleP1.MirrorCipher("aabcc")

	println("🚪:EscapeTheLockedDoor: doorCombination is ", escapeTheLockedDoor)
	println("🪞:MirrorCipher: unique symbol count is: ", mirrorCipher)

}
