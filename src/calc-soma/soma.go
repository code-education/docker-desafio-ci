package main

import "fmt"

func soma(x int, y int) int {
	return x + y;
}

func main() {
	var x,y = 5, 5
	resultado := soma(x, y)
	fmt.Println("A soma de ", x, " + ", y, " = ", resultado);
}