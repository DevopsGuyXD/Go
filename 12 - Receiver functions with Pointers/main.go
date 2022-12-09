package main

import "fmt"

func main() {

	myBill := newBill("Marios bill")

	myBill.addItem("Onion soup", 4.50)
	myBill.addItem("Veg pie", 3.75)
	myBill.addItem("Toffee pudding", 5.67)
	myBill.addItem("Coffee", 1.50)

	myBill.updateTip(10)

	fmt.Println(myBill.format())
}