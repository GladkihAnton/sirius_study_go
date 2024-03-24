package structs

import "fmt"

type Address struct {
	street string
}

func (address *Address) getAddress() {
	fmt.Printf("old address: %v", address.street)
}

type User struct {
	a int
	Address
}

func (user *User) setUser() {
	user.a = 10
}

func (user *User) getAddress() {
	fmt.Printf("New address %v", user.street)
	user.Address.getAddress()
}

func (user *User) setUserPointer() {
	user.a = 10
}

func (user *User) getUser() {
	fmt.Printf("user: %v", user.street)
	fmt.Println()
	user.getAddress()
}

func getUser(user *User) {
	fmt.Printf("user: %v", user.street)
	fmt.Println()
	user.getAddress()
}

func Main() {
	user := User{a: 1, Address: Address{street: "asda"}}
	user.getUser()
	user.setUser()
	user.getUser()
	user.setUserPointer()
	user.getUser()
}
