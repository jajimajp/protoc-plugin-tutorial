package proto

type Person struct {
	Name    string
	Age     int32
	Company Company
}
type Company struct {
	Name string
}
