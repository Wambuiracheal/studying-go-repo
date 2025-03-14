// GENERATE RANDOM INT
// package main

// import(
// 	"fmt"
// 	"math/rand"
// )

// func main(){
// 	fmt.Println("My favorite number is: ", rand.Intn(10) )
// }

// SQUAREROOT ROUNDED OFF TO THE NEAREST
// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main (){
// 	fmt.Printf("Now you have %g problems. \n", math.Sqrt(7))
// 	fmt.Println(math.Pi)
// }

// // PI
// package main

// import(
// 	"fmt"
// 	"math"
// )

// func main(){
// 	fmt.Println(math.Pi)
// }

// // ADDITION
// package main

// import "fmt"

// func add(x int, y int) int {
// 	return x + y 
// }

// func main(){
// 	fmt.Println(add(23, 45))
// }

// // SWAP FUNCTION
// package main

// import "fmt"

// func swap(x, y string)(string, string){
// 	return y, x
// }

// func main(){
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// }

// SPLIT
package main

import "fmt"

func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

func main(){
	fmt.Println(split(17))
}

