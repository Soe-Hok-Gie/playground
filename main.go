package main

import (
	"fmt"
	algoritma "playground/algoritma"
)

func main() {

	transactions := []algoritma.Transaction{
		{Id: 1, User: "Anas", Amount: 50},
		{Id: 2, User: "Muin", Amount: 60},
		{Id: 3, User: "Nawar", Amount: 70},
	}

	//FilterAndSort
	filtered := algoritma.FilterAndSort(transactions, "Nawar")
	fmt.Println(filtered)

	//TotalPerUser
	Totals := algoritma.TotalPerUser(transactions)
	fmt.Println(Totals)
}
