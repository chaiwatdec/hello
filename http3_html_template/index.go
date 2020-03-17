package main
import (
//	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

type Product struct{
	Name string
	Price int
}

func main() {
	var templates = template.Must(template.ParseFiles("index.html"))
	router:=mux.NewRouter()
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		Myproduct:=Product{"milk",500}	//add product detail
		templates.ExecuteTemplate(w, "index.html",Myproduct)
	})	//call back function

	http.ListenAndServe(":8080",router)
}