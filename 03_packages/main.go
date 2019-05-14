package main

import (
	"fmt"
	"math"

	"github.com/maheshuligade/go_crash_course/04_functions/strutil"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
	fmt.Println(greeting("myname"))
	fmt.Println(getSum(3, 4))
}
