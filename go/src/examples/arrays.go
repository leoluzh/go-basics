package main 

import (
	"fmt"
)

func main(){
	//array with fixed dimension
	var a [5] int 
	fmt.Println("Print array a : " , a )
	a[1] = 5
	a[3] = 7
	a[4] = 9
	fmt.Println("Print array a again but values are setted with index",a)

	//initialization of array ... using type inference again
	b := [5]int{1,2,3,4,5}
	fmt.Println("Showing values of array b " , b )

	//no fixed bounded array - slice - like List/ArrayList in Java
	c := []int{2,4,6,8,10}
	fmt.Println("Showing slice c values: " , c )

	//append a value into slice ... use build func append
	c = append( c , 12 )
	fmt.Println("Showing slice c values with appened new value ", c ) 		
	array := []string{ "a" , "b" , "c" , "d" , "e"  } 
	for index , value := range array {
		fmt.Println(" index and value " , index , value )
	}

	m := make( map[string]string ) 
	m["a"] = "alpha"
	m["b"] = "beta"
	
	for key , value := range m {
		fmt.Println("key:" , key , " value: " , value  )
	}

}
