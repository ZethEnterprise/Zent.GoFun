package main

import
(
	"fmt"
	"html"
	"log"
	"net/http"
        
        "github.com/go-openapi/loads"
        "github.com/go-openapi/runtime/middleware"
        
        "go-rest-api/pkg/swagger/restapi"
        "go-rest-api/operations"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})

	log.Println("Listening on loaclhost:8080")

	log.Fatal(http.ListenAndServe(":8080",nil))
}