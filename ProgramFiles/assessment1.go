package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now()
	fmt.Println("\nToday : ", loc, " Time : ", now.String())


	ChrismusDate := time.Date(2021, time.December, 25, 12, 10, 52, 211, time.UTC)
	fmt.Println("ChrismusDate  : ", loc, " Time : ", ChrismusDate) //
	diff := ChrismusDate.Sub(now)

	hrs := int(diff.Hours())
	fmt.Printf("Diffrence in Hours : %d Hours\n", hrs)

	mins := int(diff.Minutes())
	fmt.Printf("Diffrence in Minutes : %d Minutes\n", mins)

	second := int(diff.Seconds())
	fmt.Printf("Diffrence in Seconds : %d Seconds\n", second)

	days := int(diff.Hours() / 24)
	fmt.Printf("Diffrence in days : %d days\n", days)

}

