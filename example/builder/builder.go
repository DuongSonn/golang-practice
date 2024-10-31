package main

type Student struct {
	Name   string
	Age    int
	Gender string
}

type StudentRegisterInterface interface {
	SetName(s string) StudentRegisterInterface
	SetAge(n int) StudentRegisterInterface
	SetGender(s string) StudentRegisterInterface
	Register() Student
}

type StudentRegister struct {
	student Student
}

func NewStudentRegister() StudentRegisterInterface {
	return &StudentRegister{}
}
func (r *StudentRegister) SetName(s string) StudentRegisterInterface {
	r.student.Name = s
	return r
}
func (r *StudentRegister) SetAge(n int) StudentRegisterInterface {
	r.student.Age = n
	return r
}
func (r *StudentRegister) SetGender(s string) StudentRegisterInterface {
	r.student.Gender = s
	return r
}
func (r *StudentRegister) Register() Student {
	return r.student
}
