package main

import (
	"fmt"
	"errors"
)


// how can this `"jibreel"` be equivalent to something like `{Name: "jibreel", Password: "12345678"}`
type User struct{
	Name string
	Password string
}


const Signup int = 1
const Login int = 2
const Exit int = 3



var users = make(map[string]string)

// var users = []User{}

func main() {
	loop:
	for {
		choice := chooseOption()
		
		switch choice {
		case Signup:
			signup()
		case Login: 
			login()
		case Exit:
			fmt.Println("You are exiting. Take care!")
			break loop;
		default:
			fmt.Println("Invalid choice")
		}
	}
}


func chooseOption() int {
	fmt.Println("1. Signup")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")
	var choice int
	fmt.Scanln(&choice)
	return choice
}


func signup() {
	user := getUser()

	exists := userExists(user.Name)
	if exists {
		fmt.Println("A user with this name already exists! Try another name.")
		return;
	}
	users[user.Name] = user.Password
	fmt.Printf("welcome %v\n", user.Name)
}

func login() {
	user := getUser()
	err := authenticate(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Hoorey! you have sucessfully logged in as %v\n", user.Name)
}


func getUser() User {
	fmt.Println("Enter your name: ")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter your password: ")
	var password string
	fmt.Scanln(&password)

	return User{Name: name, Password: password}
}

func userExists(name string) bool {
	_, exists := users[name]
	return exists 
}


func authenticate(user User) error {
	password, ok := users[user.Name]
	if !ok {
		return errors.New("User does not exist")
	}
	if user.Password == password {
		return nil
	}
	return errors.New("wrong password")
}
