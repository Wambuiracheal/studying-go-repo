// package main

// import "fmt"

// func main() {
//     var a int = 58  // Declare an integer variable `a`
//     var ptr *int    // Declare a pointer to an integer

//     ptr = &a // Assign the memory address of `a` to `ptr`

//     fmt.Println("Value of a:", a)        // Prints the value of `a`
//     fmt.Println("Address of a:", &a)     // Prints the memory address of `a`
//     fmt.Println("Pointer value (address):", ptr)  // Prints the memory address stored in `ptr`
//     fmt.Println("Dereferenced pointer value:", *ptr) // Dereference `ptr` to get the value of `a`
// }


// CHANGING VALUES USING POINTERS
// package main

// import "fmt"

// func main() {
//     x := 10
//     fmt.Println("Before:", x)

//     // Pass pointer of x to the function
//     modifyValue(&x)

//     fmt.Println("After:", x) // The value of x is changed to 20
// }

// func modifyValue(ptr *int) {
//     *ptr = 20 // Dereference the pointer and modify the value at that memory address
// }

// POINTERS AND ARRAYS
package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println("Before:", arr)

    // Passing the pointer to the first element of the array
    modifyArray(&arr)

    fmt.Println("After:", arr) // The array values are modified
}

func modifyArray(arr *[3]int) {
    arr[0] = 10 // Modify the first element of the array using the pointer
}

