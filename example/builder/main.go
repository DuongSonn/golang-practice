package main

import "fmt"

func main() {
	studentRegister := NewStudentRegister()
	eleStudent := studentRegister.SetAge(6).SetName("A").SetGender("male").Register()
	fmt.Println("Elementary Student: ", eleStudent)
}
