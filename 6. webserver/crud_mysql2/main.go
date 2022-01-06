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
	FUNCTION ALL STUDENT
*/
func all(){
	db, err := connectDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var age = 27
	rows, err := db.Query("SELECT id, name, grade FROM tb_student where age = ? ",age)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []student

	for rows.Next(){
		var each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil{
			fmt.Println(err.Error())
			return
		}

		result = append(result,each)

	}

	if err = rows.Err(); err != nil{
		fmt.Println(err.Error())
		return
	}

	for _, each := range result{
		fmt.Println(each.name)
	}
}

/*
	FUNCTION BY ID
*/
func findById(id string){
	var db, err = connectDB()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	var result = student{}
	err = db.QueryRow("SELECT name, grade FROM tb_student WHERE id = ? ", id).Scan(&result.name, &result.grade)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}

func main(){
	//find all
	all()
	//find by id
	findById("B001")


}
