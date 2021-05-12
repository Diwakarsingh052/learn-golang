package main

import "fmt"

type FullTimeEmp struct {
	empId int
	name  string
	emp   EmpInfo // embed the EmpInfo struct // It will give access to all of it's fields
}
type ContractEmp struct {
	name string
	emp  EmpInfo // field name will be struct name
}

type EmpInfo struct {
	WorkingHours int
	pay          float64
}

func main() {

	ft := FullTimeEmp{
		empId: 101,
		name:  "Rajesh",
		emp: EmpInfo{
			WorkingHours: 8,
			pay:          20000,
		},
	}

	fmt.Println(ft.emp.WorkingHours)

}
