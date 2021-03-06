//Kieran O'Halloran

package main

//importing sort package and fmt
import (
	"fmt"
	"sort"
)

func merge() {

	var a [5]int
	//ask user for list of values
	fmt.Println("Please enter 5 numbers for 1st array")

	//populate array with values given
	for i := 0; i < 5; i++ {
		//ask user for array values
		fmt.Println("Please enter number: ")

		fmt.Scanf("%d ", &a[i])
	}
	var b [5]int
	//ask user for values
	fmt.Println("Please enter 5 numbers for 2nd array")

	//populate array with values
	for i := 0; i < 5; i++ {
		//ask user for array values
		fmt.Println("Please enter number: ")
		fmt.Scanf("%d ", &b[i])
	}

	x, y := a[:5], b[:5]
	fmt.Printf("x: %v, y: %v\n", x, y)

	//join two lists
	x = append(x, y...)

	sort.Ints(x)
	//prints sorted list
	fmt.Printf("Sorted list: %v", x)
}
