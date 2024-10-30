package main

import "fmt"

func main() {
	makeConcurrentRequestV1()
	// makeConcurrentRequestV2()

	studentRegister := NewStudentRegister()
	eleStudent := studentRegister.SetAge(6).SetName("A").SetGender("male").Register()
	fmt.Println("Elementary Student: ", eleStudent)

	customLog()
}
