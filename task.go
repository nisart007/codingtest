package main

import "fmt"

//PrintNumbers from 1-10
func PrintNumbers() {
	fmt.Println("")
	fmt.Printf("%d", 1)
	for i := 2; i <= 10; i++ {
		fmt.Printf(",%d", i)
	}
	fmt.Println("")
}

//AreaOfTraingle Area Of Traingle
func AreaOfTraingle(w, h float64) float64 {
	area := (w * h) / 2
	fmt.Printf("Area of Traingle, width: %f, height: %f is :%f\n", w, h, area)
	return area
}
