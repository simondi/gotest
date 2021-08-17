package main

// This is an import block
import  (
    "fmt"
)

/*
The main functions
*/
func main() {
    

    // Array variable
    arr := [3]int{1,2,3};
    fmt.Println(arr);

	arr1 := []int{1,2,3};
    fmt.Println(arr1);

    arr1 = append(arr1, 4, 5, 6);
    fmt.Println(arr1);

}