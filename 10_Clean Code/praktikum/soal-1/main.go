package main

import "fmt"

//Penamaan struct dan properties harus dimulai dengan Uppercase
type User struct {
	Id       int
	Username string // Username dan Password harus menggunakan tipe data string
	Password string
}

type UserService struct { //Penamaan harus menggunakan camel case
	Users []User //Penamaan harus jelas
}

//Menambahkan method Add User
func (u *UserService) AddUser(newUser User) {
	u.Users = append(u.Users, newUser)
}

//Penamaan method harus camel case dan receiver harusnya diakses melalui pointer
func (u *UserService) GetAllUsers() []User {
	return u.Users
}

//Penamaan method harus camel case dan receiver harusnya diakses melalui pointer
func (u *UserService) GetUserById(id int) User {
	for _, r := range u.Users {
		if id == r.Id {
			return r
		}
	}

	return User{}
}

//Menambahkan func main
func main() {
	// Membuat instance
	userService := UserService{}

	// Menambah User
	user1 := User{Id: 1, Username: "nurulalyh", Password: "lulu"}
	user2 := User{Id: 2, Username: "tiara_juli", Password: "yaya"}
	user3 := User{Id: 3, Username: "aisyah_rp", Password: "anggi"}

	userService.AddUser(user1)
	userService.AddUser(user2)
	userService.AddUser(user3)

	// Menampilkan semua data user
	Users := userService.GetAllUsers()
	fmt.Println("Data All User : ")
	for _, val := range Users {
		fmt.Println("Id : ", val.Id)
		fmt.Println("Username : ", val.Username)
		fmt.Println("Password : ", val.Password)
		fmt.Println("")
	}

	// Menampilkan user berdasarkan Id
	UserId := 2
	userById := userService.GetUserById(UserId)
	if userById.Id != 0 {
		fmt.Println("Data ditemukan!")
		fmt.Println("Id : ", userById.Id)
		fmt.Println("Username : ", userById.Username)
		fmt.Println("Password : ", userById.Password)
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}
