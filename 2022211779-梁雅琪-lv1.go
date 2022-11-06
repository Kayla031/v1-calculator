package main

import "fmt"

func main() {
	calc()
}
func calc() {
	var a, b float32
	var cmd string
	var c []float32
	for {
		fmt.Scan(&a, &cmd, &b)
		switch cmd {
		case "+":
			c = append(c, a+b)
			fmt.Println(c)
		case "-":
			c = append(c, a-b)
			fmt.Println(c)
		case "*":
			c = append(c, a*b)
			fmt.Println(c)
		case "/":
			c = append(c, a/b)
			fmt.Println(c)
		}
	}
}
