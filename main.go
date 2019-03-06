package main

import "log"

// Sum is function for adding two ints
func Sum(x, y int) int {
	return x + y
}

//Sub is
func Sub(x, y int) (result int) {
	result = x - y
	return
}

func main() {
	log.Println("Main starting....")
	log.Println("Sum(5, 6) = ", Sum(5, 6))
	log.Println("Sub(5, 6) = ", Sub(5, 6))
}
