package main

import (
	"fmt"
	"github.com/arifbugaresa/mnc-logic/tasks"
	"time"
)

func DisplayInfo(number int) {
	fmt.Println("\n=====================")
	fmt.Println(" Logic Test", number)
	fmt.Println("=====================\n")
}

func main() {
	Task1()
	Task2()
	Task3()
	Task4()
}

// Task1 logic 1
func Task1() {
	DisplayInfo(1)

	n := 4
	stringsList := []string{
		"abcd",
		"acbd",
		"aaab",
		"acbd",
	}
	fmt.Println(tasks.FindMatch(n, stringsList))
}

// Task2 logic 2
func Task2() {
	DisplayInfo(2)

	totalBill := 700649
	amountPaid := 800000

	fmt.Println(tasks.CalculateChange(totalBill, amountPaid))

	totalBill = 575650
	amountPaid = 580000

	fmt.Println(tasks.CalculateChange(totalBill, amountPaid))

	totalBill = 657560
	amountPaid = 600000

	fmt.Println(tasks.CalculateChange(totalBill, amountPaid))
}

// Task3 Logic 3
func Task3() {
	DisplayInfo(3)

	tasks.IsValidBracketString("[>")
}

// Task4 Logic4
func Task4() {
	DisplayInfo(4)

	var (
		totalHoliday  float64 = 7
		leaveDuration float64 = 1
		joinDate, _           = time.Parse("2006-01-02", "2021-05-01")
		leaveDate, _          = time.Parse("2006-01-02", "2021-07-05")
	)

	canTake, reason := tasks.LeaveRequest(joinDate, leaveDate, totalHoliday, leaveDuration)

	fmt.Printf("Can take leave: %v\n", canTake)
	fmt.Printf("Reason: %s\n", reason)
}
