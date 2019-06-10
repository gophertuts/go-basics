package main

import "fmt"

// Person represents the person struct
type Person struct {
	Name string
	Age  int
}

// People slice
type People []Person

// BestPeople array
type BestPeople [2]Person

// Grades map
type Grades map[string]int

// Learner interface
type Learner interface {
	Learn()
}

// IsMathExpert func
type IsMathExpert func(Grades) bool

func main() {
	// Also called Composite Literals

	p1 := Person{Name: "Steve", Age: 24}
	p2 := Person{"John", 19}
	p3 := struct {
		Name string
		Age  int
	}{
		Name: "Judy",
		Age:  22,
	}
	people1 := People{p1, p2, p3}
	people2 := []Person{p1, p2, p3}
	bestPeople := BestPeople{p1, p2}
	//bestPeople := BestPeople{p1, p2, p3} // Index out of bounds
	grades1 := map[string]int{
		"Math":    10,
		"English": 9,
		"Physics": 9,
	}
	grades2 := Grades{
		"Math":    10,
		"English": 9,
		"Physics": 9,
	}
	grades3 := grades1
	// Reference type (it will affect grades3 as well)
	// The copy of variable points to the same address memory
	grades1["Math"] = 8

	isMathExpert := func(g Grades) bool {
		if v, ok := g["Math"]; ok && v == 10 {
			return true
		}
		return false
	}

	fmt.Printf("Person:\n%+v\n%+v\n%+v\n\n", p1, p2, p3)
	fmt.Printf("People:\n%+v\n%+v\n\n", people1, people2)
	fmt.Printf("Best people:\n%+v\n\n", bestPeople)
	fmt.Printf("Grades:\n%+v\n%+v\n%+v\n\n", grades1, grades2, grades3)
	fmt.Printf("IsMathExpert:\n%+v\n%+v\n\n", isMathExpert(grades1), isMathExpert(grades2))
}
