package main

func main() {

	//for loop using slice 
	slice := []int{1, 2, 3};
	for j, v := range slice {
		println(j, v);
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v);
	}
}