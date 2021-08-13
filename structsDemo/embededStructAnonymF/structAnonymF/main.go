package main

//
//type Circle struct {
//	X,Y,Radius int
//}
//type Wheel struct {
//	X,Y, Radius, Spokes int
//}

type Point struct {
	X, Y int
}
type Circle struct {
	center Point
	Radius int
}
type Wheel struct{
	Circle Circle
	Spokes int
}

var w Wheel

func main() {
	w.Circle.center.X = 8  // equivalent to w.Circle.Point.X = 8
	w.Circle.center.Y = 8  //equivalent to w.Circle.Point.Y = 8
	w.Circle.Radius = 5			//equivalent to w.Circle.Radius = 5
	w.Spokes   = 20

	//As the set of shapes grows, we’re bound to notice similarities and rep etition among them, so it
	//may be convenient to fac tor out their common par ts:

}

