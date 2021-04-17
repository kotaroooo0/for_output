package main

import "github.com/kotaroooo0/for_output/protobuf/tutorialpb"

func main() {
	p := tutorialpb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*tutorialpb.Person_PhoneNumber{
			{Number: "555-4321", Type: tutorialpb.Person_HOME},
		},
	}
}
