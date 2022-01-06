/* 
	Type 1 
*/

package main

import "fmt"

type Employee struct{
	name string
	address string
}

type Grade struct{
	name string
	Employee
}

//nested stuct
type student struct {
    person struct {
        name string
        age  int
    }
    grade   int
    hobbies []string
}

// tag property dalam struct
type person struct {
    name string `tag1`
    age  int    `tag2`
}

func main(){

	/* =========================== STRUCT =================================== */

	// type 1
	var emp = Grade{}
	emp.name = "Hisbikal"
	emp.address = "Kembang Lio"
	emp.Employee.name = "A"
	fmt.Println(emp.name)
    fmt.Println(emp.address)
    fmt.Println(emp.Employee.name)


	// type 2
	var emp1 = Employee{name:"luna bella",address:"Bella Cassa"}
	var emp2 = Grade{name: "B",Employee: emp1}
	fmt.Println(emp1.name)
    fmt.Println(emp1.address)
    fmt.Println(emp2.name)

	//type 3 kombinasi slice dengan struct
	var employeemii = []Employee{
		{name: "Jack",address:"Kembang Lio"},
		{name: "Dwi",address:"Bella Cass"},
	}

	for _, emp := range employeemii{
		fmt.Println("name : ",emp.name," address: ",emp.address)
	}

	//type 4 struct with alias
	type People = Employee
	var p1 = People{}
	p1.name = "People Bikal"
	fmt.Println(p1.name)

	/* ========================================================================== */

	/* ====================== ANONYMOUS STRUCT ================================== */

	//type 1 tanpa pengisian property
	var empDepart = struct {
		Employee
		depart string
	}{}
	empDepart.Employee.name = "Hisbikal"
	empDepart.depart = "IT PROGRAMMER"
	fmt.Println(empDepart.Employee.name)
    fmt.Println(empDepart.depart)

	//type 2 dengan pengisian property
	var empSite = struct{
		Employee
		site string
	}{
		Employee: Employee{"Bella","Kembang Lio"},
		site: "BRI RAGUNAN",
	}
	fmt.Println(empSite.Employee.name)
    fmt.Println(empSite.site)

	//type inisialisasi slice dengan struct anonymous
	var allStudents = []struct {
		Employee
		grade int
	}{
		{Employee: Employee{"wick", "Kembang Lio"}, grade: 2},
		{Employee: Employee{"ethan", "Kembang Lio"}, grade: 3},
		{Employee: Employee{"bond", "Kembang Lio"}, grade: 3},
	}
	
	for _, student := range allStudents {
		fmt.Println(student)
	}
	/* ====================== ANONYMOUS STRUCT ================================== */


}