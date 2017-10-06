//Kieran O'Halloran

//Adapted from : https://play.golang.org/p/vhHmjhOMEo

package main

import "fmt"

func main() { //function to create a list of numbers
	x := []int{
		8, 9, 65, 87,
		57, 22, 52, 44,
		47, 94, 83, 27,
		19, 97, 89, 17,
	}

	smallest, biggest := x[0], x[0] //set largest and smallest to 0
	for _, v := range x {           //loop through the list
		if v > biggest { //if the number is bigger than the last sorted largest overwrite it
			biggest = v
		}
		if v < smallest { //if the number is smaller than the last sorted smallest overwrite it
			smallest = v
		}
	}
	//print out largest and smallest
	fmt.Println("The biggest number is ", biggest)
	fmt.Println("The smallest number is ", smallest)
}
