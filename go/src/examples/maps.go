package main 

import(
	"fmt"
)

func main(){
	vertices := make( map[string]int )
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["pentagon"] = 5
	vertices["dodecagon"] = 12
	
	fmt.Println("vertices: " , vertices )
	fmt.Println("vertices[square] ", vertices["square"] )

	//remove element form the map
	fmt.Println("Using buildin function delete to remove element...")
	delete( vertices , "pentagon" )
	fmt.Println("vertices: " , vertices )
	

}
