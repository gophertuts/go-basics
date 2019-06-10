package main

import "fmt"

type student struct {
	grade float32
	name  string
}

type teacher struct {
	grade float32
	name  string
}

func (s *student) learn() {
	fmt.Println("student learns")
}

func (t *teacher) teach() {
	fmt.Println("teacher teaches")
}

func main() {
	s := student{
		grade: 9,
		name:  "John",
	}
	t := teacher{
		grade: 10,
		name:  "Alex",
	}
	// Types should be comparable
	promotedStudent := teacher(s)

	fmt.Printf("Student: %+v", s)
	fmt.Printf("Teacher: %+v", t)
	fmt.Printf("Converted student: %+v", promotedStudent)
}
