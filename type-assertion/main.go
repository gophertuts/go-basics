package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Human represents an abstract Human type (any kind of human)
type Human interface {
	Money()
}

// Student represents student type
type Student struct {
	Name   string
	Grades map[string]int
}

// Average represents student's average grade
func (s *Student) Average() float64 {
	var sum int
	for _, g := range s.Grades {
		sum += g
	}
	return float64(sum / len(s.Grades))
}

// Money represents the amount of money a student makes
func (s *Student) Money() {
	fmt.Printf("%s's grades: %v\n", s.Name, s.Grades)
	fmt.Printf("%s's salary is: $%f/day\n", s.Name, s.Average())
}

// Teacher represents the teacher type
type Teacher struct {
	Name   string
	Salary float64
}

// Money represents the amount of money a teacher makes
func (t *Teacher) Money() {
	fmt.Printf("%s's teacher salary is: $%f/day | $%f/month", t.Name, t.Salary/30, t.Salary)
}

// Teach represents the teacher's main skill
func (t *Teacher) Teach(subject string, students ...Student) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	fmt.Println(len(students))
	for _, s := range students {
		t.Salary += 50
		s.Grades[subject] = r.Intn(10)
	}
}

func main() {
	s1 := Student{Name: "John", Grades: make(map[string]int)}
	s2 := Student{Name: "Alex", Grades: make(map[string]int)}
	s3 := Student{Name: "Mike", Grades: make(map[string]int)}
	t := Teacher{Name: "Simon"}

	t.Teach("Math", s1, s2, s3)
	t.Teach("Physics", s1, s2)
	t.Teach("English", s1, s3)
	t.Teach("Biology", s2, s3)

	s1.Money()
	s2.Money()
	s3.Money()
	t.Money()
}
