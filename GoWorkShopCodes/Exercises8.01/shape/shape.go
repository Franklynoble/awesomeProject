package main


type Shape interface {
	area() float64
	name() string
}
type Triangle struct {
	Base   float64
	Height float64
}
type Rectangle struct {
	Length float64
	Width  float64
}
type Square struct {
	Side float64
}
