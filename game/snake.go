package game

type direction int
type Snake struct {
	body      []int     // Snake body
	length    int       // Snake Length
	direction direction // Direction snake is facing
}

// Got help from https://programming.guide/go/define-enumeration-string.html and https://golangbyexample.com/iota-in-golang/
const (
	UP direction = iota
	DOWN
	LEFT
	RIGHT
)
