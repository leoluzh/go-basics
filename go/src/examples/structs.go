package main 

import(
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) print() {
	fmt.Println("Name: " , p.Name , " Age: " , p.Age )
}

func main(){
	p := Person{ Name: "Jon Doe" , Age: 31  }
	fmt.Println("Person" , p )
	p.print()
}

