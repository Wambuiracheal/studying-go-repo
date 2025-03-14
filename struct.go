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

Ways to Initialize a Struct
Using a Struct Literal (Field Names Specified)

// way 1
person := Person{Name: "Bob", Age: 30}
Without Field Names (Order Matters)

// way 2
person := Person{"Charlie", 35} // Not recommended (reduces readability)
Using new (Pointer to Struct)

// way 3
personPtr := new(Person)
personPtr.Name = "David"
personPtr.Age = 40
fmt.Println(personPtr.Name, personPtr.Age)
Using Struct Pointers Explicitly

// way 4
person := &Person{Name: "Eve", Age: 28}
fmt.Println(person.Name, person.Age) // No need to dereference
Struct Methods
Structs can have methods by defining functions with a receiver:

// way 5
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}

func main() {
    person := Person{Name: "Alice", Age: 25}
    person.Greet() // Output: Hello, my name is Alice
}

// Structs with Pointers (Modifying Struct Fields)
// By default, Go passes struct copies to functions/methods. To modify the original struct, use a pointer:

// // 6
func (p *Person) HaveBirthday() {
    p.Age++ // Modifies the original struct
}

func main() {
    person := Person{Name: "Alice", Age: 25}
    person.HaveBirthday()
    fmt.Println(person.Age) // Output: 26
}

// Anonymous Structs
// For quick usage without defining a type:

person := struct {
    Name string
    Age  int
}{"Frank", 40}

fmt.Println(person.Name, person.Age)
Embedded Structs (Composition)

// Go does not have inheritance, but you can embed structs to achieve similar behavior.

type Address struct {
    City, Country string
}

type Employee struct {
    Name    string
    Age     int
    Address // Embedded struct
}

func main() {
    emp := Employee{Name: "Grace", Age: 32, Address: Address{City: "New York", Country: "USA"}}
    fmt.Println(emp.Name, emp.City, emp.Country) // Direct access to embedded fields
}

// Tags in Structs (Useful for JSON, Database, etc.)

// Go supports struct tags, which provide metadata for serialization:

import "encoding/json"

type User struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

func main() {
    user := User{Name: "John", Email: "john@example.com"}
    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData)) // Output: {"name":"John","email":"john@example.com"}
}


// Key Takeaways
// ✅ Structs help organize related data into a single unit.
// ✅ They support methods to encapsulate behavior.
// ✅ Use pointers to modify struct fields inside methods.
// ✅ Embedded structs allow for a composition-based design.
// ✅ Struct tags help with encoding/decoding data formats like JSON.

// Would you like an example of struct usage in a real-world Go application? 🚀
