package main

import (
	"fmt"
	"net/http"

	"github.com/SkylerBair/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Printf("starting server at :3000")
	http.ListenAndServe(":3000", nil)
}
