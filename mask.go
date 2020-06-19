package main

import "math"

type maskFunc func(x, y int) bool

var masks map[int]maskFunc = map[int]maskFunc{
	0: func(x, y int) bool { return (x+y)%2 == 0 },
	1: func(x, y int) bool { return (x)%2 == 0 },
	2: func(x, y int) bool { return (y)%3 == 0 },
	3: func(x, y int) bool { return (x+y)%3 == 0 },
	4: func(x, y int) bool { return int(math.Floor(float64(x)/2)+math.Floor(float64(y)/3))%2 == 0 },
	5: func(x, y int) bool { return ((x*y)%2)+((x*y)%3) == 0 },
	6: func(x, y int) bool { return (((x*y)%2)+((x*y)%3))%2 == 0 },
	7: func(x, y int) bool { return (((x+y)%2)+((x*y)%3))%2 == 0 },
}
