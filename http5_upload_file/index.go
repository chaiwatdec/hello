package main
import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/upload",uploadHandle)

	http.ListenAndServe(":8080",nil)
}

func index(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w,r,"upload.html")
}

func uploadHandle(w http.ResponseWriter, r *http.Request){
	file,handle,err:=r.FormFile("file") //Get request Form name:file //handle is map
	defer file.Close()
	if err!=nil {
		fmt.Println(err)
		return
	}
	//fmt.Fprintf(w, "%v",handle.Header) //print on web
	fileup,err:=os.OpenFile("./test/"+handle.Filename,os.O_CREATE,0666) //refer pkg os
	if err!=nil {
		fmt.Println(err)
		return
	}
	defer fileup.Close()
	io.Copy(fileup,file)	//refer pkg os
	fmt.Fprintf(w,"Upload Complete")
}