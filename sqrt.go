// learning how to round square root without using the function "sqrt".

package main

import (
	"fmt"
)

func sqrt(x float64) string {
	if x < 0 { 
		return sqrt(-x) + "i"		// for complex numbers
    } 
  	var z float64 = 1				  // here you can change the "z" number that goes on in the first interaction 
  	for i:=0; i< 5; i++ {			// you may change "i" numbers of interaction also
	  	z -= (z * z - x) / (2 * z)}
  return fmt.Sprint(z)        // still don't know why to use this, but i know i need to
}


func main() {
	fmt.Println(sqrt(9))
}
