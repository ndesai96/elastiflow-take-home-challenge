package main

import (
	"fmt"

	"github.com/ndesai96/elastiflow-take-home-challenge/tree-search/tree"
)

func main() {
	tests := []*tree.Tree{
		buildEmployeeGraphWithDuplicates(true),
		buildEmployeeGraphWithDuplicates(false),
	}

	for _, employeeGraph := range tests {
		employee, level := tree.CheckDuplicateIDs(employeeGraph)
		if employee != nil {
			fmt.Printf("Employee ID: %d, Level: %d\n", employee.GetID(), level)
		} else {
			fmt.Println("No duplicate employee!")
		}
	}
}

type Employee struct {
	id int
}

func (t *Employee) GetID() int {
	return t.id
}

func buildEmployeeGraphWithDuplicates(withDuplicates bool) *tree.Tree {
	ceo := tree.NewNode(&Employee{id: 1})

	employee2 := tree.NewNode(&Employee{id: 2})
	employee3 := tree.NewNode(&Employee{id: 3})
	employee4 := tree.NewNode(&Employee{id: 4})
	employee5 := tree.NewNode(&Employee{id: 5})
	employee6 := tree.NewNode(&Employee{id: 6})
	employee7 := tree.NewNode(&Employee{id: 7})
	employee8 := tree.NewNode(&Employee{id: 8})

	ceo.AddChildren(employee6, employee8)
	employee6.AddChildren(employee2, employee5, employee7)
	employee5.AddChildren(employee4, employee3)

	if withDuplicates {
		employee5.AddChildren(employee7)
		employee3.AddChildren(employee8)
	}

	return tree.New(ceo)
}
