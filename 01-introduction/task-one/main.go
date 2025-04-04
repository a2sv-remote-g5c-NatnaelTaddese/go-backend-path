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

func calculateAverage() {
	average = total / float32(n_subjects)
	fmt.Print(Blue)
	fmt.Println("------------------------------------")
	fmt.Printf("Your average grade is %.2f\n", average)
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
		fmt.Print(Cyan, " Subject ")
		fmt.Scanln(&subject)
		if subject != "" {
			return subject
		}
		fmt.Println(Red, "Invalid subject. Please enter a valid subject.", Reset)
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
	print("How many subjects have you taken? ")
	fmt.Scanln(&n_subjects)

	for i := 0; i < n_subjects; i++ {
		subject := getAndValidateSubject()
		grade := getAndValidateGrade(subject)
		grades[subject] = grade
		total += grade
	}

	calculateAverage()
}
