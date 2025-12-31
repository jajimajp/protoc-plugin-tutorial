package proto

import (
	fmt "fmt"
	strings "strings"
)

type Person struct {
	Name    string
	Age     int32
	Company Company
}

func (m Person) PrintYAML(indentSpaces int) {
	indent := strings.Repeat(" ", indentSpaces)
	fmt.Printf("%sName: %v\n", indent, m.Name)
	fmt.Printf("%sAge: %v\n", indent, m.Age)
	fmt.Printf("%sCompany: \n", indent)
	m.Company.PrintYAML(indentSpaces + 2)
}

type Company struct {
	Name string
}

func (m Company) PrintYAML(indentSpaces int) {
	indent := strings.Repeat(" ", indentSpaces)
	fmt.Printf("%sName: %v\n", indent, m.Name)
}
