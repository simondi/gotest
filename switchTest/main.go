package main

type HTTPRequest struct {
	Method string;
}

func main() {
	r := HttpRequest{Method: "GET"};
	
	switch r.Method {
	case "GET":
		println("GET Request")
	case "DELETE":
		println("DELETE Request")
	case "POST":
		println("POST Request")
	case "PUT":
		println("PUT Request")
	default:
		println("Unhandled")
	}
}