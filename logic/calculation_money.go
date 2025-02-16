package logic

import "fmt"

func CalculateChange(totalBill, amountPaid int) string {
	fmt.Println("Total belanja seorang customer: Rp.", totalBill)
	fmt.Println("Pembelil membayar: Rp.", totalBill)

	// Check if the amount paid is less than the total bill
	if amountPaid < totalBill {
		return "False, kurang bayar"
	}

	// Calculate the change to be given
	change := amountPaid - totalBill

	// Round down the change to the nearest multiple of 100
	change = (change / 100) * 100

	// Available denominations
	denominations := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}
	changeDetails := make(map[int]int)

	// Calculate the number of bills and coins for each denomination
	for _, denom := range denominations {
		denomCount := change / denom
		if denomCount > 0 {
			changeDetails[denom] = denomCount
			change %= denom
		}
	}

	// Format the output details for the change and denominations
	result := fmt.Sprintf("Kembalian yang harus diberikan kasir: %d, dibulatkan menjadi %d \n\nPecahan uang:\n", amountPaid-totalBill, amountPaid-totalBill-((amountPaid-totalBill)%100))

	for _, denom := range denominations {
		if count, exists := changeDetails[denom]; exists {
			// Display as "bills" for amounts >= 1000 and "coins" for amounts < 1000
			if denom >= 1000 {
				result += fmt.Sprintf("%d Lembar %d\n", count, denom)
			} else {
				result += fmt.Sprintf("%d Koin %d\n", count, denom)
			}
		}
	}

	return result
}
