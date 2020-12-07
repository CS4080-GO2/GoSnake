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
	fieldHeight = 20
	WallColor   = termbox.ColorCyan
	EmptyColor  = termbox.ColorDefault
	SnakeColor  = termbox.ColorRed
)

var width int
var height int

func InitField() Field {
	f := Field {
		food: 	InitFood(),
		snake: 	InitSnake(fieldWidth, fieldHeight),
		height: fieldHeight,
		width: 	fieldWidth,
	}

	return f
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
	width, height = termbox.Size()

	// Make bottom
	for x := 0; x < fieldWidth + 1; x++ {
		termbox.SetCell(x, fieldHeight, ' ', WallColor, WallColor)
	}

	// Make top
	for x := 0; x < fieldWidth + 1; x++ {
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

// Function for snake movement
func (f *Field) move() {
	head := f.snake.body[0]

	switch f.snake.direction {
	case UP:
		f.snake.moveBody(Coordinate{x: head.x, y: head.y - 1})
	case DOWN:
		f.snake.moveBody(Coordinate{x: head.x, y: head.y + 1})
	case LEFT:
		f.snake.moveBody(Coordinate{x: head.x - 1, y: head.y})
	case RIGHT:
		f.snake.moveBody(Coordinate{x: head.x + 1, y: head.y})
	}

	// Check if the snake hit its body


	// If the snake exit the field then display "Game Over"
	f.SnakeExit()
}

// Function for when the snake leaves the field
func (f * Field) SnakeExit() {
	head := f.snake.body[0]

	if head.x >= fieldWidth || head.y >= fieldHeight ||
		head.x <= 0 || head.y <= 0 {
		// If the leaves the field, it's game over
		GameOver()
	}
}
