package main
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//connect db [db, err := sql.Open("mysql", "user:password@/dbname")]
	db, err := sql.Open("mysql", "root:@/pythondb")
	if err!=nil{
		fmt.Println(err)
	}
	defer db.Close()	//close connection

	//update data
	stmt,err:=db.Prepare("update tbl_member set member_age=? where member_id=?")
	stmt.Exec(60,"6")

	if err!=nil{
		fmt.Println(err)
	}

	//query data
	rows,err:=db.Query("select * from tbl_member order by member_id asc")
	if err!=nil{
		fmt.Println(err)
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
