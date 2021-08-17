package main

// This is an import block
import  (
	"net/http"
	"github.com/pluralsight/webservice/controllers"
    //"fmt"

	//"github.com/pluralsight/webservice/models"
)

/*
The main functions
*/
func main() {
    controllers.RegisterController();
	http.ListenAndServe(":3000", nil);
}