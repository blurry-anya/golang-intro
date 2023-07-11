package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	newBill := newBill(name)

	fmt.Println("Bill is created!")

	return newBill
}

func promptOptions(bill bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a - add an item, s - save bill, t - add a tip) ", reader)
	switch option {
	case "a":
		fmt.Println("Adding a new item..")

		name, _ := getInput("Name: ", reader)
		price, _ := getInput("Price: ", reader)

		floatPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("A price must be a number! \n Returning...")
			promptOptions(bill)
		}

		bill.addItem(name, floatPrice)

		fmt.Printf("Added an item %q with a price %0.2f\n", name, floatPrice)
		promptOptions(bill)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		floatTip, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("A tip must be a number! \n Returning...")
			promptOptions(bill)
		}
		bill.setTip(floatTip)

		fmt.Printf("Added a tip %0.2f to the %q\n", bill.tip, bill.name)
		promptOptions(bill)
	case "s":
		fmt.Println("Saving the bill..")
		bill.save()
	default:
		fmt.Println("Not a valid option!")
		promptOptions(bill)
	}

}

func structsAndCustomTypes() {
	myBill := createBill()
	promptOptions(myBill)
}
