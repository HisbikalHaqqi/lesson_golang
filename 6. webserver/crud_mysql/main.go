/*
	SQL QUERY TANPA PREPARE STATEMENT
*/
package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

/*
	STRUCT DATA
 */
type student struct {
	id    string
	name  string
	age   int
	grade int
}


/*
	CONNECT MYSQL
 */
func connectDB()(*sql.DB, error){
	db, err := sql.Open("mysql","root:@tcp(127.0.0.1)/go")

	if(err != nil){
		return nil, err
	}

	return db,nil
}

/*
	FUNCTION WITH PREPARE STATAMENT
 */
func queryPrepare(id *string) {
	db, err := connectDB()

	if err != nil{
		fmt.Println(err.Error())
		return
	}

	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = student{}
	stmt.QueryRow(id).Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)
}

/*
	FUNCTION CRUD WITH PREPARED STATEMENT
 */
func crudWithPrepare(){
	db, error := connectDB()

	if error != nil{
		fmt.Println(error.Error())
		return
	}
	defer db.Close()

	//INSERT WITH PREPARE STATEMENT
	var stmt, err = db.Prepare("INSERT INTO tb_student VALUE (?,?,?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	stmt.Exec("G001", "Galahad", 29, 2)

	//UPDATE WITH PREPARE STATEMENT
	var stmt2, err2 = db.Prepare("UPDATE tb_student SET name = ? , age = ? , grade = ? WHERE id = ?")
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}
	stmt2.Exec("Hisbikal", 23, 1, "W001")

	//DELETE WITH PREPARE STATEMENT
	var stmt3, err3 = db.Prepare("DELETE FROM tb_student WHERE id = ?")
	if err3 != nil {
		fmt.Println(err.Error())
		return
	}
	stmt3.Exec( "B002")
}

func sqlExec() {
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G002", "Galahad 2", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G002")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")
}

func main(){
	//find by id
	var id string
	fmt.Println("pilih id")
	fmt.Scanln(&id)
	queryPrepare(&id)

	//crud with prepare statement
	crudWithPrepare()

	//crud with exec
	sqlExec()
}
