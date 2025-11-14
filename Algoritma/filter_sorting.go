package algoritma

import "sort"

type Transaction struct {
	Id     int
	User   string
	Amount int
}

func FilterAndSort(transaction []Transaction, user string) []Transaction {
	//step 1
	var result []Transaction
	for _, t := range transaction {
		if t.User == user {
			result = append(result, t)
		}

	}
	//step 2 : Sort DESC
	sort.Slice(result, func(i, j int) bool {
		return result[i].Amount > result[j].Amount
	})
	return result

}
