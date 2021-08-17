package main

// This is an import block
import  (
    "fmt"
)

const (
    pi = 3.14159;
    landa = 5.6;
    a = iota + 4;
    b= iota;
    c;
)
/*
The main functions
*/
func main() {
    // print a variable of integer
    fmt.Println("Hello, Simon");
    var i int = 42;
    fmt.Println("Hello, ", i);

    // string variable declaration 1:
    firstName := "Arthur";
    fmt.Println("Hello, ", firstName );

    // string variable declaration 1:
    var lastName  string = "McDonald";
    fmt.Println("Hello, ", lastName);

    //Pointer
    var nameList *string = new(string);
    *nameList = "Arthur";
    fmt.Println(*nameList);

    //pointer use reference address
    ptr  := &firstName;
    fmt.Println(ptr, *ptr);

    //iota
    fmt.Println(a, b, c);

}