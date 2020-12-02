package game

import (
	"github.com/nsf/termbox-go"
)

type Field struct {
	food   Food
	snake  Snake
	height int
	width  int
}

const (
	fieldWidth = 50
	fieldHeight = 30
	wallColor = termbox.ColorCyan
	emptyColor = termbox.ColorWhite
	snakeColor = termbox.ColorRed
)

var width int
var height int

// Display draws the field, snake, and food.
func Display() {
	// Clear screen.
	termbox.Clear(emptyColor,emptyColor)
	width, height = termbox.Size()
	// Make border
	makeBorder()
	makeSnake()


}

func makeSnake() {

}

func makeBorder() {
	// Make top
	for x := 1; x < fieldWidth; x++ {
		termbox.SetCell(x, fieldHeight, ' ', wallColor, wallColor)
	}
	// Make bottom
	for x := 1; x < fieldWidth; x++ {
		termbox.SetCell(x, 0, ' ', wallColor, wallColor)
	}
	// Make right
	for y := 1; y < fieldHeight; y++ {
		termbox.SetCell(fieldWidth, y, ' ', wallColor, wallColor)
	}
	// Make left
	for y := 1; y < fieldHeight; y++ {
		termbox.SetCell(0, y, ' ', wallColor, wallColor)
	}

	// Now display it
	termbox.Flush()
}
