package main

import  (
		"fmt"
		"flag"
		"strconv"
		"math"
		)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
    fmt.Println("Hello, world")
	fmt.Println(Sqrt(4))
	
	for i, v := range flag.Args() {
		fmt.Println(i, v)
		f, error := strconv.Atof64(v)
		if error != nil {
		  fmt.Println("err: " + error.String())
			return
		}
		fmt.Println(Sqrt(f))

	}
	
	
}
