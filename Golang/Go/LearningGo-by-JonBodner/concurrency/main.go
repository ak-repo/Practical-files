package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	x int
	y int
}

func main() {

	p := Point{20, 20}

	x := "hii"
	fmt.Println(reflect.TypeOf(p).Kind())
	fmt.Println(reflect.TypeOf(x))

	var pi *int= &p.x
	fmt.Println("pi",*pi)

}
