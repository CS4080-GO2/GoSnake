package game

import (
	"math/rand"
)
// Thanks to https://unicode-table.com/en/#supplemental-symbols-and-pictographs
var foodList = []rune{
	'🥓',
	'🥔',
	'🥛',
	'🥑',
}

type Food struct {
	char	rune
	coord	Coordinate
}

func InitFood() Food {
	return Food {
		char: foodList[rand.Intn(len(foodList)-1)],
		// TODO make coord that isn't on snake.
	}
}


func DropFoodAt(c Coordinate) Food {
	return Food{
		char:	foodList[rand.Intn(len(foodList))],
		coord:	c,
	}
}
