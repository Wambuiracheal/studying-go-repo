package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // Creating a struct instance
    person1 := Person{Name: "Alice", Age: 25}
    
    // Accessing struct fields
    fmt.Println("Name:", person1.Name)
    fmt.Println("Age:", person1.Age)
}
