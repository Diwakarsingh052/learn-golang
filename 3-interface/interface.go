package main

import "fmt"

type Expense interface {
	CalculateSalary() int
}

type perm struct {
	empId    int
	basicPay int
	pf       int
}

type contract struct {
	empId    int
	basicPay int
	bonus    int
}

func (p perm) CalculateSalary() int {
	return p.basicPay + p.pf

}
func (c contract) CalculateSalary() int {
	return c.basicPay + c.bonus
}
func (c contract) Show() { // not part of interface signature // can only be called with struct var
	fmt.Println(c.empId)
}

func totalExpense(e ...Expense) { // this will accept any struct instance that implements the interface


	total := 0


	for _, v := range e {

		total = total + v.CalculateSalary()
	}
	fmt.Println("Total Expense", total)
}

func main() {

	e1 := perm{
		empId:    101,
		basicPay: 20000,
		pf:       1000,
	}

	e2 := contract{
		empId:    102,
		basicPay: 25000,
		bonus:    15000,
	}
	e3 := perm{
		empId:    103,
		basicPay: 22000,
		pf:       18000,
	}

	totalExpense(e1, e2, e3)

}
