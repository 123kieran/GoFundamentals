//Kieran O'Halloran

package main

import "fmt"

func fizz() {

	for i := 1; i <= 100; i++ {
		//if number can be divided by 3
		if i%3 == 0 {
			//if number can be divided by 15 (3*5)
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
			//if number can be divided by 5
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
