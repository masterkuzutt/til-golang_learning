package main

import (
	"fmt"
	_ "til-golang_learning/gin-sample/work/imp"
	"time"
)

type User interface {
	PrintDetails()
	PrintName()
}

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

func (p Person) PrintName() {
	fmt.Println(p.FirstName, p.LastName)
}
func (p Person) PrintDetails() {
	fmt.Println(p.Email, p.Location, p.Dob.String())

}

type Admin struct {
	Person
	Roles []string
}

func (a Admin) PrintDetails() {
	a.Person.PrintDetails()
	for _, v := range a.Roles {
		fmt.Println(v)
	}
}

func main() {
	hoge := Admin{
		Person{
			"Takahiro",
			"Sakai",
			time.Now(),
			"hoge@com",
			"tokyo",
		},
		[]string{"admin", "worker"},
	}
	hoge.PrintDetails()
	hoge.PrintName()

	user := []User{hoge}

	for _, v := range user {
		v.PrintDetails()
	}

}
