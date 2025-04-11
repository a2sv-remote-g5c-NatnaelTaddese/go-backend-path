package main

import (
	"fmt"
	"strings"
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
	userName       string
	grades         map[string]float32
	nSubjects      int
	total, average float32
)

func getGradeColor(grade float32) string {
	switch {
	case grade < 50:
		return Red
	case grade < 80:
		return Yellow
	default:
		return Green
	}
}

func calculateAverage() {
	average = total / float32(nSubjects)
	fmt.Print(Blue)
	fmt.Println("------------------------------------")
	fmt.Print("Your average grade is")
	fmt.Print(getGradeColor(average))
	fmt.Printf(" %.2f\n", average)
	fmt.Println("------------------------------------")
	fmt.Print(Reset)
}

func getAndValidateGrade(subject string) float32 {
	var grade float32
	for {
		fmt.Printf("%s Grade for %s: %s", Yellow, subject, Cyan)
		_, err := fmt.Scanln(&grade)
		if err == nil && grade >= 0 && grade <= 100 {
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
		subject = strings.TrimSpace(subject)
		if subject != "" {
			if _, exists := grades[subject]; exists {
				fmt.Println(Red, "This subject already exists. Please enter a different subject.", Reset)
				continue
			}
			return subject
		}
		fmt.Println(Red, "Invalid subject. Please enter a valid subject name.", Reset)
	}
}

func getAndValidateNumberOfSubjects() int {
	var n int
	for {
		fmt.Print(Purple, "How many subjects have you taken? ", Cyan)
		_, err := fmt.Scanln(&n)
		if err == nil && n > 0 {
			return n
		}
		fmt.Println(Red, "Invalid input. Please enter a positive number of subjects.", Reset)
	}
}

func displayGradeTable() {
	fmt.Print(Blue)
	fmt.Println("------------------------------------")
	fmt.Printf("%-15s | %-10s\n", "SUBJECT", "GRADE")
	fmt.Println("------------------------------------")

	for subject, grade := range grades {
		fmt.Printf("%s%-15s | %s%-10.2f%s\n", Blue, subject, getGradeColor(grade), grade, Blue)
	}
}

func getValidName() string {
	var name string
	for {
		fmt.Print(Purple, "Please Enter your Name: ", Cyan)
		fmt.Scanln(&name)
		name = strings.TrimSpace(name)
		if name != "" {
			return name
		}
		fmt.Println(Red, "Invalid name. Please enter a valid name.", Reset)
	}
}

func init() {
	fmt.Println(Green, "Initializing Grade Calculator...", Reset)
	grades = make(map[string]float32)
}

func main() {
	userName = getValidName()
	fmt.Println(Purple, "Welcome", userName, "to the Grade Calculator!", Reset)
	nSubjects = getAndValidateNumberOfSubjects()

	fmt.Println(Yellow, "Please enter your subjects and grades:", Reset)
	for i := range nSubjects {
		fmt.Printf(Blue+"\nSubject %d of %d:"+Reset+"\n", i+1, nSubjects)
		subject := getAndValidateSubject()
		grade := getAndValidateGrade(subject)
		grades[subject] = grade
		total += grade
	}

	fmt.Println(Green, "\nGrade Summary for", userName, Reset)
	displayGradeTable()
	calculateAverage()

	fmt.Print(Purple)
	if average >= 90 {
		fmt.Println("Excellent work! You're in the top tier.")
	} else if average >= 80 {
		fmt.Println("Very good! Keep up the great work.")
	} else if average >= 70 {
		fmt.Println("Good job! You're doing well.")
	} else if average >= 60 {
		fmt.Println("Satisfactory. There's room for improvement.")
	} else if average >= 50 {
		fmt.Println("You passed, but consider studying more next time.")
	} else {
		fmt.Println("You need to improve your grades. Consider seeking academic help.")
	}
	fmt.Print(Reset)
}
