package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, err:=sql.Open("mysql", "<username>:<password>@/pythondb")
	db, err := sql.Open("mysql", "root:@/pythondb")
	if err!=nil{
		fmt.Println(err)	//print error
	}
	defer db.Close()	//close connection

	//update data
	stmt,err:=db.Prepare("update tbl_member set member_age=? where member_id=?")	//delete data if column:member_id=?:7
	stmt.Exec(700,"7")

	if err!=nil{
		fmt.Println(err)	//print error
	}

	//query data
	rows,err:=db.Query("select * from tbl_member order by member_id asc")
	if err!=nil{
		fmt.Println(err)	//print error
	}
	for rows.Next(){
		var id string
		var name string
		var lname string
		var age int
		err=rows.Scan(&id,&name,&lname,&age)

		fmt.Printf("ID: %s Name: %s LName: %s Age: %d\n",id,name,lname,age)
	}
	
}