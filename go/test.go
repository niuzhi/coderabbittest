package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func getUser(id int) *User {
	if id <= 0 {
		return nil
	}
	return &User{Name: fmt.Sprintf("User%d", id), Age: 20 + id}
}

func main() {
	user := getUser(-1)
	fmt.Println("用户名：", user.Name)
}
