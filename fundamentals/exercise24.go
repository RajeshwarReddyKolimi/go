package main

import (
	"fmt"
	"sort"
)

type School struct {
	roster map[int][]string
}

func New() *School {
	return &School{roster: make(map[int][]string)}
}

func (s *School) addStudent(name string, grade int) {
	s.roster[grade] = append(s.roster[grade], name)
}

func (s *School) getStudentsOfAGrade(grade int) []string {
	return s.roster[grade]
}

func (s *School) getAllStudentsSorted() []string {
	allStudents := []string{}
	grades := []int{}
	for g := range s.roster {
		grades = append(grades, g)
	}
	sort.Ints(grades)
	for _, grade := range grades {
		sort.Strings(s.roster[grade])
		allStudents = append(allStudents, s.roster[grade]...)
	}
	return allStudents
}

func (s *School) getAllStudentsWithGrades() map[int][]string {
	return s.roster
}

func (s *School) getGradeOfAStudent(student string) int {
	for grade, students := range s.roster {
		for _, name := range students {
			if student == name {
				return grade
			}
		}
	}
	return -1
}

func Exercise24() {
	school := New()
	school.addStudent("Nemo", 1)
	school.addStudent("Jim", 2)
	school.addStudent("Dora", 2)
	school.addStudent("Mini", 3)
	g1 := school.getStudentsOfAGrade(1)
	g2 := school.getStudentsOfAGrade(2)
	g3 := school.getStudentsOfAGrade(3)
	all := school.getAllStudentsWithGrades()
	sorted := school.getAllStudentsSorted()
	jim := school.getGradeOfAStudent("Jim")
	kim := school.getGradeOfAStudent("Kim")
	fmt.Println("Grade 1:", g1)
	fmt.Println("Grade 2:", g2)
	fmt.Println("Grade 3:", g3)
	fmt.Println("All students:", all)
	fmt.Println("All students sorted:", sorted)
	printGrade(jim, "Jim")
	printGrade(kim, "Kim")
}

func printGrade(grade int, name string) {
	if grade != -1 {
		fmt.Println(name, "grade:", grade)
	} else {
		fmt.Println(name, "is not present")
	}
}
