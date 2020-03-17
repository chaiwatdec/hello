package main
import (
//	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"html/template"
)

type product struct{
	name string
	price int
}

func main() {
	var templates = template.Must(template.ParseFiles("index.html"))
	router:=mux.NewRouter()
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		myproduct:=product{"milk",500}	//add product detail
		templates.ExecuteTemplate(w, "index.html",myproduct)
	})	//call back function

	http.ListenAndServe(":8080",router)
}