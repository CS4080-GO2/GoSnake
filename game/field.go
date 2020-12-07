package game

import (
	"os"
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
	fieldHeight = 22
	WallColor   = termbox.ColorCyan
	EmptyColor  = termbox.ColorDefault
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

func (f *Field) move() {
	head := f.snake.body[0]
	c := Coordinate{x: head.x, y: head.y}

	// switch s.direction { // Current direction
	// case UP:
	// 	if termbox.GetCell(head.x, head.y-1).Bg == EmptyColor { // TODO change to pick up food.
	// 		s.moveBody(Coordinate{x: head.x, y: head.y - 1})
	// 	} else {
	// 		// Collision
	// 	}
	// case DOWN:
	// 	if termbox.GetCell(head.x, head.y+1).Bg == EmptyColor { // TODO change to pick up food.
	// 		s.moveBody(Coordinate{x: head.x, y: head.y + 1})
	// 	} else {
	// 		// Collision
	// 	}
	// case LEFT:
	// 	if termbox.GetCell(head.x-1, head.y).Bg == EmptyColor { // TODO change to pick up food.
	// 		s.moveBody(Coordinate{x: head.x - 1, y: head.y})
	// 	} else {
	// 		// Collision
	// 	}
	// case RIGHT:
	// 	if termbox.GetCell(head.x+1, head.y).Bg == EmptyColor { // TODO change to pick up food.
	// 		s.moveBody(Coordinate{x: head.x + 1, y: head.y})
	// 	} else {
	// 		// Collision
	// 	}
	// }

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

	// Check if the head is on the body
	if f.snake.headOnBody(c) {
		os.Exit(0)
	}

	f.SnakeExit()

}

// Function for when the snake leaves the field
func (f * Field) SnakeExit() {
	head := f.snake.body[0]

	if head.x >= fieldWidth || head.y >= fieldHeight ||
		head.x <= 0 || head.y <= 0 {
		GameOver()
	}
}
