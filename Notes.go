package main

import fmt "fmt"

type Foo int

type Vertex struct {
	X int
	Y int
}

func add(x int, y int) int {
	return x + y
}


func main() {
	fmt.Printf("Hello, world; or Καλημέρα κόσμε; or こんにちは 世界\n")	
	fmt.Println(add(1, 3))
	
	var f Foo = 1
	g := Foo(2)
	fmt.Println(g)
	fmt.Println(f)
	
	// no pointer arithmetic
	
	var  v Vertex
	q := &v
	
	// Maps = HashMap-like
	// map[string]int
	//var m map[string]int	
	//m = make(map[string]int)
	// or just do
	m := make(map[string]int)
	
	m["Bootcamp"] = 2011
	i := m["Bootcamp"]
	
	
	var n = map[string]bool {
		"true": true,
		"false": false,
	}
	
	type Cache map[string]*Entry
	
	o := Cache{"x": &Entry{1}, "y": &Entry{2}}
	
	// Slices - like smart arrays	
	// []int
	// []*Entry
	
	var x []*Entry // len(x) == 0
	y := make([]*Entry, 10) // len(y) == 10
	z := make([]*Entry, 0, 10) // len(z) == 0, cap(z) == 10  - capacity
	
	y[0] = 10
	a := y[0] // == 10
	
	b := z[0] // index overflow, 0 >= len(z)
	
	
}
