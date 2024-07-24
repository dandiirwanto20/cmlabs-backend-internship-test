package helpers

import "fmt"

func GetFizzBuzzLogic(startLoop int, endLoop int) {

	for i := startLoop; i < endLoop; i++ {
		// jika keliapatan 3 dan 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "FizzBuzz")
			continue
		}
		// jika kelipatan 5
		if i%5 == 0 {
			fmt.Println(i, "Buzz")
			continue
		}
		// jika kelipatan 3
		if i%3 == 0 {
			fmt.Println(i, "Fizz")
			continue
		}
		// Jika angka bukan kelipatan 3, 5 atau 3 dan 5
		// fmt.Printf("Angka %d bukan merupakan Fizz, Buzz, ataupun FizzBuzz \n", i)
	}
}
