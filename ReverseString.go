//Kieran O'Halloran
//https://coderwall.com/p/fw6miw/reverse-text-in-golang

package main

import "fmt"

func Reverse(s string) string {
	var result string
	//looping from end of string to element 0 and populating result with characters in each element
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}

//main function
func main() {
	//declare variable for input string
	var str string
	//ask user for string
	fmt.Println("Please enter string: ")
	//scans user input into console
	fmt.Scanf("%s ", &str)
	fmt.Println()
	//outputs result
	fmt.Println("Reverse string: ", Reverse(str))
}
