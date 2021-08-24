package main

// This is an import block
import  (
    "fmt"
    "regexp"
    "strings"
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
    var greeting string = "Hello, Simon";
    fmt.Println(greeting);

    r, _ := regexp.Compile(`S([a-z]+)n`);
    greeting = r.ReplaceAllString(greeting,"Apple");
    fmt.Println(greeting);

    greeting = strings.ToLower(greeting);
    fmt.Println(greeting);

    greeting = strings.Title(greeting);
    fmt.Println(greeting);


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