package main

import "fmt"

type User struct {
	id int64
	name string
	surename string
	login string
	password string
	gender string
	age int64
	userinfo string
}

func changeinfo(person User, info string){
	person.userinfo = info
	fmt.Println(person)
}
func main() {
	NewUser := User{
		id:       12345678,
		name:     "Ahmadjon",
		surename: "Rajabov",
		login:    "Ahmad06",
		password: "25794613",
		gender:   "M",
		age:      16,
		userinfo: "Ya-teenager",
	}
	//var a int64 = 10
	//a := int64(10)
	//var a int64
	//a=10
	fmt.Println(NewUser)
	//NewUser.userinfo="Adult"
	//fmt.Println(NewUser)
	changeinfo(NewUser,  "Adult")
}
