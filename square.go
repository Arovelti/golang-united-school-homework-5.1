package square

// package main

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	// implement me
	return Point{
		x: s.start.x + int(s.a),
		y: s.start.y + int(s.a),
	}
}

func (s *Square) Area() uint {
	// implement me
	return s.a * s.a
}

func (s *Square) Perimeter() uint {
	// implement me
	return s.a * 4
}

// func main() {
// 	sq := &Square{
// 		Point{
// 			x: 10,
// 			y: 20,
// 		},
// 		30,
// 	}
// 	fmt.Println(sq)

// 	sq1 := sq.End()
// 	sq2 := sq.Area()
// 	sq3 := sq.Perimeter()

// 	fmt.Println(sq1, sq2, sq3)
// }
