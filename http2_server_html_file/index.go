package main
import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	router:=mux.NewRouter()
	testDB:=map[string]int{"test1":10,"test2":20,"test3":30,}
	
	router.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"index.html")
	})	//call back function

	router.HandleFunc("/signup",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"signup.html")
	})	//call back function

	router.HandleFunc("/login",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"login.html")
	})	//call back function

	router.HandleFunc("/file",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"name.txt")
		//	fmt.Fprintf(w,"TuM")
	})	//call back function

	router.HandleFunc("/test/{name}",func(w http.ResponseWriter, r *http.Request){
		vars:=mux.Vars(r)
		name:=vars["name"]
		num:=testDB[name]
		fmt.Fprintf(w,"Test: %s %d",name,num)
	}).Methods("GET")

	http.ListenAndServe(":8080",router)
}