package main

import "fmt"

func main() {
	intslice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range intslice {
		if intslice[i]%2 == 0 {
			fmt.Println(intslice[i], " is even!")
		} else {
			fmt.Println(intslice[i], " is odd!!")

		}
	}

}
