package main

import (
	"fmt"

	"github.com/maheshuligade/go_crash_course/03_packages/strutil"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {

	fmt.Println(strutil.Reverse("olleh"))
	fmt.Println(greeting("myname"))
	fmt.Println(getSum(3, 4))
}
