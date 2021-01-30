package main 

import (
	"fmt"
)

func main(){
	var x int
	fmt.Println("var x int - without initalization " , x )
	x = 5 
	fmt.Println("var x receive value after initialization", x )
	
	var y int = 7

	//type inference
	sum := x + y
 
	fmt.Println("x + y " , sum )


	var i = 0
	fmt.Println("var i = " , i )

}
