package main

// This is an import block
import  (
    "fmt"
	"github.com/pluralsight/webservice/models"
)

/*
The main functions
*/
func main() {
    u := models.User{
		ID: 2,
		FirstName: "Simon",
		LastName: "Mcmillab",
	}

	fmt.Println(u)

}