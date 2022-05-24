package main

import "fmt"

func main() {
	prime_seed := [10]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	for element := range prime_seed {
		fmt.Println(prime_seed[element])
		for n := 2; n < 100; n++ {
			fmt.Println((prime_seed[element] * n) + (n - 1))
		}
	}
}
