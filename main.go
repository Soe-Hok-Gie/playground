package main

import (
	"fmt"
	algoritma "playground/algoritma"
)

func main() {

	transactions := []algoritma.Transaction{
		{Id: 1, User: "Ahmad", Amount: 50},
		{Id: 2, User: "Awanda", Amount: 60},
		{Id: 3, User: "Jek", Amount: 70},
	}
	//FilterAndSort
	filtered := algoritma.FilterAndSort(transactions, "Jek")
	fmt.Println(filtered)

	//TotalPerUser
	Totals := algoritma.TotalPerUser(transactions)
	fmt.Println(Totals)
}
