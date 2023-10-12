package main

import "fmt"

type Point struct {
	X, Y float64
}

func (p Point) Add(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}
func (p Point) Sub(q Point) Point {
	return Point{
		X: p.X + q.X,
		Y: p.Y + q.Y,
	}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {

	b := make(Path, 2)
	b = []Point{
		{1, 2},
		{2, 4},
	}
	a := Point{
		X: 1,
		Y: 1,
	}
	b.TranslateBy(a, true)
	fmt.Println(b)
	fmt.Println(a)

}
