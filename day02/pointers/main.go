package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func reset(v Vertex) {
	// set v.X and v.Y to 0
}

func resetPtr(v *Vertex) {
	// set v.X and v.Y to 0
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	v2 := v
	v2.Y = 99
	fmt.Println(v, v2)

	reset(v)
	fmt.Println(v)

	resetPtr(&v)
	fmt.Println(v)
}
