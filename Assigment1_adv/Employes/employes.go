package main

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}
type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("ID: %d, Name: %s, Salary: %d тг", fte.ID, fte.Name, fte.Salary)
}

func (pte PartTimeEmployee) GetDetails() string {
	totalEarnings := float32(pte.HourlyRate) * pte.HoursWorked
	return fmt.Sprintf("ID: %d, Name: %s, Hourly Rate: %d тг, Hours Worked: %.2f, Total Earnings: %.2f тг",
		pte.ID, pte.Name, pte.HourlyRate, pte.HoursWorked, totalEarnings)
}

type Company struct {
	Collective map[string]Employee
}

func (c Company) AddEmployee(emp Employee) {
	var empID string

	switch e := emp.(type) {
	case FullTimeEmployee:
		empID = fmt.Sprint(e.ID)
	case PartTimeEmployee:
		empID = fmt.Sprint(e.ID)
	}

	_, exists := c.Collective[empID]
	if exists == true {
		fmt.Println("Employee already exist")
	}
	c.Collective[empID] = emp
}

func (c Company) ListEmployees() {
	for _, empl := range c.Collective {
		fmt.Println(empl.GetDetails())
	}
}
func main() {

	company := Company{Collective: make(map[string]Employee)}

	company.ListEmployees()
}
