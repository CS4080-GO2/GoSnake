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
	fieldWidth  = 50
	fieldHeight = 30
	WallColor   = termbox.ColorCyan
	EmptyColor  = termbox.ColorWhite
	SnakeColor  = termbox.ColorRed
)

var width int
var height int

func InitField() Field {
	return Field {
		food: InitFood(),
		snake: InitSnake(fieldWidth, fieldHeight),
		height: fieldHeight,
		width: fieldWidth,
	}
}

// Display draws the field, snake, and food.
func (f *Field) Display() {
	// Clear screen.
	termbox.Clear(EmptyColor, EmptyColor)

	// Make border
	drawBorder()
	drawSnake(&f.snake)

	// Now display it
	termbox.Flush()
}

func drawSnake(s *Snake) {
	for i := 0; i < len(s.body); i++ {
		termbox.SetCell(s.body[i].x, s.body[i].y, ' ', SnakeColor, SnakeColor)
	}
}

func drawBorder() {
	// Make bottom
	for x := 1; x < fieldWidth; x++ {
		termbox.SetCell(x, fieldHeight, ' ', WallColor, WallColor)
	}
	// Make top
	for x := 1; x < fieldWidth; x++ {
		termbox.SetCell(x, 0, ' ', WallColor, WallColor)
	}
	// Make right
	for y := 1; y < fieldHeight; y++ {
		termbox.SetCell(fieldWidth, y, ' ', WallColor, WallColor)
	}
	// Make left
	for y := 1; y < fieldHeight; y++ {
		termbox.SetCell(0, y, ' ', WallColor, WallColor)
	}
}
