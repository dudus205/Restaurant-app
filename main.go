package main

import (
	"fmt"
)

func createBill() bill {
	var name string
	fmt.Print("Create a new bill name: ")
	fmt.Scanln(&name)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	var opt string
	fmt.Print("Choose option (a -add item, s - save bill, t - add tip): ")
	fmt.Scanln(&opt)

	switch opt {
	case "a":
		var name string
		var price float64

		fmt.Println("What is a name of an item: ")
		fmt.Scanln(&name)
		fmt.Println("What is a price of an item: ")
		fmt.Scanln(&price)

		b.addItem(name, price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("File name is: " + b.name)
		promptOptions(b)
	case "t":
		var tip float64

		fmt.Println("Tip amount ($): ")
		fmt.Scanln(&tip)

		b.updateTip(tip)
		promptOptions(b)
	default:
		fmt.Println("That was not right option...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()

	promptOptions(myBill)

	fmt.Println("Bill name - ", myBill.name)
	fmt.Println(myBill.format())
}
