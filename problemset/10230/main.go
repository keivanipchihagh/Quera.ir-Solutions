package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	if (a + b + c == 180) && a != 0 && b != 0 && c != 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
