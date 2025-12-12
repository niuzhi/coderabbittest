package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// 获取用户信息，模拟可能返回nil的场景
func getUser(id int) *User {
	if id <= 0 {
		// 无效ID返回nil指针
		return nil
	}
	return &User{Name: fmt.Sprintf("User%d", id), Age: 20 + id}
}

func main() {
	// 传入无效ID，得到nil指针
	user := getUser(-1)
	// 缺陷：直接解引用nil指针的Name字段
	fmt.Println("用户名：", user.Name)
}
