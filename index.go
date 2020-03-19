package main
import (
	"fmt"
	"net/http"
	"time"
)

type Cookie struct{
	Name string
	Value string
	Expires time.Time		//refer package time
}

func main() {
	http.HandleFunc("/",index)
	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	expiration:=time.Now().Add(time.Hour*24*365)	//define cookie age
	cook:=http.Cookie{Name:"user",Value:"Tum",Expires:expiration} //create struct name:cook
	http.SetCookie(w,&cook)
	fmt.Fprintf(w,"Create Cookie ")
}
