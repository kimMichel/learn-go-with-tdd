package structuremethodsinterfaces

import "math"

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {

	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Width float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height / 2
}

type Form interface {
	Area() float64
}