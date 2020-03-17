package main
import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/login",login)

	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"index.html")
}

func login(w http.ResponseWriter, r *http.Request){
	fmt.Println("Method:",r.Method)	
	r.ParseForm()
	fmt.Println("UserName:",r.Form["username"]) //get value from name:username
	fmt.Println("Password:",r.Form["password"]) //get value from name:password
}
