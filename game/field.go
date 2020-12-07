package game

import (
	"math/rand"
	"time"
	"fmt"
	"github.com/nsf/termbox-go"
	rune "github.com/mattn/go-runewidth"
)

type Field struct {
	food   	Food
	snake  	Snake
	height 	int
	width	int
	points	int
}

const (
	fieldWidth  = 60
	fieldHeight = 24
	WallColor   = termbox.ColorCyan
	EmptyColor  = termbox.ColorDefault
	SnakeColor  = termbox.ColorRed
)

var width int
var height int

func InitField() Field {
	rand.Seed(time.Now().UnixNano())

	f := Field {
		snake: 	InitSnake(fieldWidth, fieldHeight),
		height: fieldHeight,
		width: 	fieldWidth,
	}

	f.PlaceFood()

	return f
}

// Display draws the field, snake, and food.
func (f *Field) Display() {
	// Clear screen.
	termbox.Clear(EmptyColor, EmptyColor)

	// Make border
	DrawBorder()
	DrawFood(f.food)
	DrawScore(f.points)
	DrawMsg(fieldWidth + 5, fieldHeight - 1, "Press ESC to exit")
	drawSnake(&f.snake)

	// Now display it
	termbox.Flush()
}

func drawSnake(s *Snake) {
	for i := 0; i < len(s.body); i++ {
		termbox.SetCell(s.body[i].x, s.body[i].y, ' ', SnakeColor, SnakeColor)
	}
}

func DrawBorder() {
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
	c := Coordinate{x: head.x, y: head.y}	// New position of the head

	switch f.snake.direction {
	case UP:
		f.snake.moveBody(Coordinate{x: head.x, y: head.y - 1})
		c.y -= 1
	case DOWN:
		f.snake.moveBody(Coordinate{x: head.x, y: head.y + 1})
		c.y += 1
	case LEFT:
		f.snake.moveBody(Coordinate{x: head.x - 1, y: head.y})
		c.x -= 1
	case RIGHT:
		f.snake.moveBody(Coordinate{x: head.x + 1, y: head.y})
		c.x += 1
	}

	// Check if the snake hit its body

	if f.snake.CheckHeadPosition(c) {	// Check if head position is on body
		// End the game, since head hit body
		GameOver("You hit your body!", f.points)
	}

	// If the snake ate the food
	if c == f.food.coord {
		go f.AddPoint(100)
		f.snake.length += 1
		f.snake.body = append(f.snake.body, c)
		f.PlaceFood()
	}

	// If the snake exit the field then display "Game Over"
	f.SnakeExit()
}

// Need to use goroutine to add the point
func (f *Field) AddPoint(point int) {
	f.points += 100
}

// Function for when the snake leaves the field
func (f *Field) SnakeExit() {
	head := f.snake.body[0]

	if head.x >= fieldWidth || head.y >= fieldHeight ||
		head.x <= 0 || head.y <= 0 {
		// If the leaves the field, it's game over
		GameOver("You're leaving the field?!", f.points)
	}
}

func DrawMsg(x, y int, msg string) {
	clr := termbox.ColorDefault

	for _, c := range msg {
		termbox.SetCell(x, y, c, clr, clr)
		x += rune.RuneWidth(c)
	}
}

func (f *Field) PlaceFood() {
	// Declare x and y coord for the rand food drop
	var randCoord Coordinate

	for {
		x := rand.Intn(fieldWidth - 2) + 1
		y := rand.Intn(fieldHeight - 2) + 1

		randCoord = Coordinate{x: x, y: y}
		if f.snake.AvailablePosition(randCoord) {
			break
		}
	}

	f.food = DropFoodAt(randCoord)
}


func DrawFood(f Food) {
	clr := termbox.ColorDefault
	termbox.SetCell(f.coord.x, f.coord.y, f.char, clr, clr)
}

func DrawScore(score int) {
	msg := fmt.Sprintf("Score: %v", score)
	DrawMsg(fieldWidth + 5, 1, msg)	// Display the score
}
