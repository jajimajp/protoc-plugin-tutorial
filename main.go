package main

import (
	"fmt"

	proto "github.com/jajimajp/protoc-plugin-tutorial/protobuf/go"
)

func main() {
	p := proto.Person{
		Name: "Taro",
		Age:  30,
		Company: proto.Company{
			Name: "FooBarCorp",
		},
	}
	fmt.Printf("%v\n", p)
}
