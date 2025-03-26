package main

import "fmt"

func createRoster(students map[string]int) map[int][]string {
	studentsGradeWise := make(map[int][]string)
	for k, v := range students {
		studentsGradeWise[v] = append(studentsGradeWise[v], k)
	}
	return studentsGradeWise
}

func getStudentsFromGrade(roster map[int][]string, grade int) []string {
	return roster[grade]
}

func Exercise5() {
	stu := map[string]int{
		"Jim":  2,
		"kim":  1,
		"Nemo": 3,
		"Dora": 2,
	}
	roster := createRoster(stu)
	grade2 := getStudentsFromGrade(roster, 2)
	fmt.Println(grade2)

}
