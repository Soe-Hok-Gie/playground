package algoritma

import (
	"sort"
)

type Transaction struct {
	Id     int
	User   string
	Amount int
}

// Tulis fungsi untuk mengambil transaksi dari user tertentu dan urutkan berdasarkan Amount DESC.
func FilterAndSort(transactions []Transaction, user string) []Transaction {
	//step 1
	var result []Transaction
	for _, t := range transactions {
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

// Tulis fungsi untuk menghitung total Amount per user.
func TotalPerUser(transactons []Transaction) map[string]int {
	totals := make(map[string]int)
	for _, t := range transactons {
		totals[t.User] += t.Amount
	}
	return totals
}
