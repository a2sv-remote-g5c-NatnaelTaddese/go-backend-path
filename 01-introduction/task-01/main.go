package main

import (
	"fmt"
)

// Color constants
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
)

var (
	user_name      string
	grades         map[string]float32
	n_subjects     int
	total, average float32
)

func getGradeColor(grade *float32) string {
	switch {
	case *grade < 49:
		return Red
	case *grade < 79:
		return Yellow
	default:
		return Green
	}
}

func calculateAverage() {
	average = total / float32(n_subjects)
	fmt.Print(Blue)
	fmt.Println("------------------------------------")
	fmt.Print("Your average grade is")
	fmt.Print(getGradeColor(&average))
	fmt.Printf(" %.2f\n", average)
	fmt.Print(Reset)
}

func getAndValidateGrade(subject string) float32 {
	var grade float32
	for {

		fmt.Print(Yellow, " Grade for ", subject, " ", Cyan)
		fmt.Scanln(&grade)
		if grade >= 0 && grade <= 100 {
			return grade
		}
		fmt.Println(Red, "Invalid grade. Please enter a grade between 0 and 100.", Reset)
	}
}

func getAndValidateSubject() string {
	var subject string
	for {
		fmt.Print(Cyan)
		fmt.Printf("%15v", "Subject: ")
		fmt.Scanln(&subject)
		if subject != "" {
			return subject
		}
		fmt.Println(Red, "Invalid subject. Please enter a valid subject.", Reset)
	}
}

func displayGradeTable() {
	fmt.Print(Blue)
	fmt.Println("------------------------------------")
	fmt.Printf("%10v : %4v \n", "SUBJECT", "GRADE")
	fmt.Print(Green)

	for subject, grade := range grades {
		fmt.Printf("%10v : %4v \n", subject, grade)

	}
}

func init() {
	fmt.Println(Green, "Initializing...", Reset)
	grades = make(map[string]float32)
}

func main() {

	fmt.Print("Please Enter your Name: ")
	fmt.Scanln(&user_name)
	fmt.Println(Purple, "Welcome", user_name, Reset)
	fmt.Print("How many subjects have you taken? ")
	fmt.Scanln(&n_subjects)

	for i := 0; i < n_subjects; i++ {
		subject := getAndValidateSubject()
		grade := getAndValidateGrade(subject)
		grades[subject] = grade
		total += grade
	}

	displayGradeTable()
	calculateAverage()
}
