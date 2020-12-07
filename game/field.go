package game

import (
	"math/rand"
	"time"
	"fmt"
	"github.com/nsf/termbox-go"
	rune "github.com/mattn/go-runewidth"
)

type Field struct {
	food   		Food			// The food.
	obstacle	Obstacle		// The obstacle
	obsList		[]Coordinate	// List of all obstacles coordinates
	snake  		Snake			// The object being controled
	height 		int				// Height of the field
	width		int				// Width of the field
	points		int				// The Score
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
var pointCap int = 500		// For bonus rounds with obstacles
var numObs int = 1

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

	// Displaying food on the field
	DrawFood(f.food)

	// Display the score
	DrawScore(f.points)

	// Msg informing players on how to exit game
	DrawMsg(fieldWidth + 5, fieldHeight - 1, "Press ESC to exit")

	// display the snake on the field
	drawSnake(&f.snake)

	// Bonus rounds!
	if f.points >= 500 {
		if f.points >= pointCap {
			f.obsList = nil

			for i := 0; i < numObs; i++ {
				// Drop the obstacle
				f.PlaceObstacle()
				f.obsList = append(f.obsList, f.obstacle.coord)
			}

			pointCap += 500
			numObs += 1
		}

		// Displaying the obstacles
		f.DrawObstacles()

		// Display New message informing player what is happening
		DrawMsg(fieldWidth + 5, fieldHeight / 2, "AVOID THE BONES!!!")
	}

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
	if c == f.food.coord ||
		((c.x == f.food.coord.x + 1) && c.y == f.food.coord.y) {
		go f.AddPoint(100)
		f.snake.length += 1
		f.snake.body = append(f.snake.body, c)
		f.PlaceFood()
	}

	if f.points >= 500 {
		if f.HitObstacle(c) {
			GameOver("Oh no! You ate the bone!", f.points)
		}
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

// Find an available space to put the food
func (f *Field) PlaceFood() {
	// Declare x and y coord for the rand food drop
	var randCoord Coordinate

	for {
		x := rand.Intn(fieldWidth - 2) + 1
		y := rand.Intn(fieldHeight - 2) + 1

		randCoord = Coordinate{x: x, y: y}
		if f.snake.AvailablePosition(randCoord) {
			if f.points >= 500 {
				if f.NotInObsPosition(randCoord) {
					break
				}
			} else {
				break
			}
		}
	}

	f.food = DropFoodAt(randCoord)
}

// Find an available space to put the obstacle
func (f *Field) PlaceObstacle() {
	var randCoord Coordinate

	for {
		x := rand.Intn(fieldWidth - 2) + 1
		y := rand.Intn(fieldHeight - 2) + 1

		randCoord = Coordinate{x: x, y: y}

		if f.snake.AvailablePosition(randCoord) {
			if randCoord != f.food.coord {
				break
			}
		}
	}

	f.obstacle = ObstacleAt(randCoord)
}


// Function to display the food on the field
func DrawFood(f Food) {
	clr := termbox.ColorDefault
	termbox.SetCell(f.coord.x, f.coord.y, f.char, clr, clr)


	// termbox.Close()
	// fmt.Println("Rune Width:  ", rune.RuneWidth(f.char))
	// fmt.Println("food x:  ", f.coord.x)
}

// Function to display the score
func DrawScore(score int) {
	msg := fmt.Sprintf("Score: %v", score)
	DrawMsg(fieldWidth + 5, 1, msg)	// Display the score
}

// Function to display the obstacles
func (f *Field) DrawObstacles() {
	clr := termbox.ColorDefault
	for i := 0; i < len(f.obsList); i++ {
		curCoord := f.obsList[i]
		termbox.SetCell(
			curCoord.x,
			curCoord.y,
			f.obstacle.char,
			clr,
			clr)
	}
}

// Function used in food drops. Makes sure that food is not in obstacle
func (f *Field) NotInObsPosition(c Coordinate) bool {
	for i := 0; i < len(f.obsList); i++ {
		if c == f.obsList[i] {
			return false
		}
	}
	return true
}

// Check if the snake ate the obstacle
func (f *Field) HitObstacle(c Coordinate) bool {
	// Check all Coordinates of the obstacles
	for i := 0; i < len(f.obsList); i++ {
		if c == f.obsList[i] {
			return true
		}
	}
	return false
}
