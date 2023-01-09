package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Employee struct {
	Id          int               `json:"id"`
	Name        map[string]string `json:"name"`
	Salary      float64           `json:"salary"`
	Designation string            `json:"designation"`
}

var err string = "Employee Record Not Found."

var employees []Employee

func input(dis string, ip interface{}) {
	fmt.Printf("Enter %s: ", dis)
	fmt.Scan(ip)
}

func showEmplyees() {
	emps, _ := json.Marshal(employees)
	fmt.Println(string(emps))
}

func showEmployee(id int) {
	for _, emp := range employees {
		if emp.Id == id {
			employee, _ := json.Marshal(emp)
			fmt.Println(string(employee))
			return
		}
	}
	panic(err)
}

func addEmployee(emp Employee) {
	employees = append(employees, emp)
	pass, _ := json.Marshal("Record Inserted Succesfully.")
	fmt.Println(string(pass))

}

func deleteEmployee(id int) {
	for index, emp := range employees {
		if emp.Id == id {
			employees = append(employees[:index], employees[index+1:]...)
			pass, _ := json.Marshal("Record Deleted Succesfully.")
			fmt.Println(string(pass))
			return
		}
	}
	panic(err)
}
func editEmployee(id int) {
	var edit int
	for index, e := range employees {
		if e.Id == id {
			fmt.Println("======Edit Options======")
			fmt.Println("1 Name")
			fmt.Println("2. Salary")
			fmt.Println("3. Designation")
			fmt.Printf("Enter Your Choice: ")
			fmt.Scan(&edit)
			switch edit {
			case 1:
				var fname, mname, lname string
				input("First Name", &fname)
				input("Middle Name", &mname)
				input("Lasst Name", &lname)
				e.Name = map[string]string{"FName": fname, "MName": mname, "LName": lname}
			case 2:
				input("Salary", &e.Salary)
			case 3:
				fmt.Printf("Enter Designation: ")
				input("Designation", &e.Designation)

			}
			employees = append(employees[:index], employees[index+1:]...)
			employees = append(employees, e)
			pass, _ := json.Marshal("Record Edited Succesfully.")
			fmt.Println(string(pass))
			return
		}
	}
	panic(err)
}

func menu(choice int) {
	switch choice {
	case 1:
		showEmplyees()
	case 2:
		var id int
		input("Employee ID", &id)
		showEmployee(id)
	case 3:
		var emp Employee
		var fname, mname, lname string
		input("First Name", &fname)
		input("Middle Name", &mname)
		input("Lasst Name", &lname)
		emp.Name = map[string]string{"FName": fname, "MName": mname, "LName": lname}
		input("Salary", &emp.Salary)
		input("Designation", &emp.Designation)
		src := rand.NewSource(time.Now().UnixNano())
		ran := rand.New(src)
		emp.Id = ran.Intn(1000)
		addEmployee(emp)

	case 4:
		var id int
		input("Employee ID", &id)
		editEmployee(id)
	case 5:
		var id int
		input("Employee ID", &id)
		deleteEmployee(id)
	}
}

func main() {
	defer showEmplyees()
	defer fmt.Println("Showing Records In DB")
	var choice int
	var check string = "y"
	for check == "y" || check == "Y" {
		fmt.Println("======MENU======")
		fmt.Println("1. Show all Employee Details.")
		fmt.Println("2. Show Employee Detail.")
		fmt.Println("3. Insert Employee Details.")
		fmt.Println("4. Update Employee Details.")
		fmt.Println("5. Delete Employee Details.")
		input("Your Choice", &choice)
		menu(choice)
		fmt.Printf("Do you want to continue:")
		fmt.Scan(&check)
	}
}
