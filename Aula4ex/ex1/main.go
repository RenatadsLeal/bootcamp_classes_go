package main

import "fmt"

type user struct {
	name     string
	surname  string
	age      int
	email    string
	password string
}

func (u *user) editName(name string) {
	u.name = name
}

func (u *user) editSurname(surname string) {
	u.surname = surname
}

func (u *user) editAge(age int) {
	u.age = age
}

func (u *user) editEmail(email string) {
	u.email = email
}

func (u *user) editPassword(password string) {
	u.password = password
}

func main() {

	user1 := user{
		name:     "Renata",
		surname:  "Leal",
		age:      33,
		email:    "renata@gmail.com",
		password: "xdhh34",
	}

	fmt.Println(user1)

	user1.editName("Renatas")
	user1.editEmail("leal@gmail.com")

	fmt.Println(user1)
	fmt.Println(user1.email)
}
