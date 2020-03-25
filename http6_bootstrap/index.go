package main
import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type ResultData struct{
	Id string
	Name string
	Lname string
	Age int
}

var templates = template.Must(template.ParseFiles("index.html"))

func main() {
	http.HandleFunc("/",resultdata)
	http.HandleFunc("/delete",deletedata)
	http.ListenAndServe(":8080",nil)
}

func deletedata(res http.ResponseWriter, req *http.Request){
	db, err := sql.Open("mysql", "root:@/pythondb")
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	stmt,err:=db.Prepare("delete from tbl_member where member_id=?")	//delete data if column:member_id=?:7
	stmt.Exec(req.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}
	http.Redirect(res,req, "/",301)
}

func resultdata(res http.ResponseWriter, req *http.Request){
	db, err := sql.Open("mysql", "root:@/pythondb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from tbl_member")
	if err != nil {
		panic(err)
	}

	tRes := ResultData{}	// set tRes as Struct:ResultData
	var results []ResultData	// var as slice get data from tRes
	for rows.Next() {	//extract data from query database per row
		var id string
		var name string
		var lname string
		var age int
		err = rows.Scan(&id,&name,&lname,&age)
		tRes.Id = id	//get data from query:id to var:tRes.Id
		tRes.Name = name
		tRes.Lname = lname
		tRes.Age = age
		results = append(results, tRes) //add data to var:results per row
		if err != nil {
			panic(err)
		}
	}
		templates.Execute(res, results)
		fmt.Println(results)
}