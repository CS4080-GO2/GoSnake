package game


var obstacleList = []rune {
    '🦴',
}


type Obstacle struct {
    char    rune
    coord   Coordinate
}

// Function that drops obstacles on the field
func ObstacleAt(c Coordinate) Obstacle {
	return Obstacle{
		char:	obstacleList[0],
		coord:	c,
	}
}
