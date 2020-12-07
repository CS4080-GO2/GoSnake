package game


type direction int

const InitialSnakeLength = 4

// Snake the actual snake
type Snake struct {
	body      []Coordinate // Snake body
	length    int          // Snake Length
	direction direction    // Direction snake is facing
}

// Got help from https://programming.guide/go/define-enumeration-string.html and https://golangbyexample.com/iota-in-golang/
const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)

func InitSnake(w, h int) Snake {
	var temp []Coordinate

	for y := 0; y < InitialSnakeLength; y++ {
		temp = append(temp, Coordinate{x: w / 2, y: (h / 2) + y}) // Essentially, we want body[0] to be the head, pointing upwards.
	}

	return Snake{
		body:      temp,
		length:    InitialSnakeLength,
		direction: UP,
	}
}

func (s *Snake) moveBody(coord Coordinate) {
	// A little jank...
	// Have to make slice without 'make' because it always allocates too much memory over and over due to the len(slice) not actually representing the number of elements in the slice.
	var temp []Coordinate
	temp = append(temp, coord)
	for i := 0; i < len(s.body)-1; i++ { // Copy all except last element over, tail needs to be forgotten.
		temp = append(temp, s.body[i])
	}

	s.body = temp
}
