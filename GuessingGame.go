//Kieran O'Halloran

//Adapted from : http://golangcookbook.blogspot.ie/2012/12/guess-number-game-v2.html

package main

import "fmt"
import "math/rand"
import "time"

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func guess() {

	var userGuess int = 0
	var count int = 0
	var exit bool = true
	randomNumber := random(1, 100)
	var numbersChosen [100]int

	for exit {

		fmt.Print("Please enter a number(1-100):\n")
		fmt.Scan(&userGuess)

		count++

		for i := 0; i < 100; i++ {
			if userGuess == numbersChosen[i] {
				fmt.Print("You already guessed that number\n")
				count--
			}
		}

		numbersChosen[count] = userGuess

		if randomNumber == userGuess {
			fmt.Printf("You guessed the number! It took you %d goes\n", count)
			exit = false
		} else if randomNumber > userGuess {
			fmt.Print("Your guess was too low, try again!\n")
		} else if randomNumber < userGuess {
			fmt.Print("Your guess was too high, try again!\n")
		}
	}

}
