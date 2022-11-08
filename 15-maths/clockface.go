package clockface

import "time"

// represents two dimensoinal cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// unit vector of the second hand of an analogue clock at time 't'
// represented as a point
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
