package main 

import(
	"fmt"
)

func main(){
	i := 7
	fmt.Println("Value of variable: " , i )
	fmt.Println("Memory address of variable" , &i )
	inc(i)
	fmt.Println("Value of variable: " ,  i )
	incr(&i)
	fmt.Println("Value of variable: " ,  i )

}

func inc( i int ) {
	i++
	fmt.Println("Passing value by value " , i )
}

func incr( i *int ){
	*i++
	fmt.Println("Passing value by reference" , *i )
}


