package fillmerchants

/*
Bolt - 1st Round
You have a sum of money to distribute among a number of merchant accounts. You need to distribute the money as evenly as possible. But some merchants have a maximum amount they can take. We only deal with dollar amounts, no decimals.

Amount: $100 --> 90 --> 70
Accounts: [('a', 20), ('b', 100), ('c', 10), ('d', 100)]

[('c', 80), ('a', 90), ('b', 100), ('d', 100)]


Result: [('a', 20), ('b', 35), ('c', 10), ('d', 35)]

Amount: $101
Accounts: [('a', 20), ('b', 100), ('c', 10), ('d', 100)]
Result: [('a', 20), ('b', 36), ('c', 10), ('d', 35)]

Amount: $200
Accounts: [('a', 20), ('b', 110), ('c', 10), ('d', 110), ('e', 100), ('f', 100)]
Result: [('a', 20), ('b', 35), ('c', 10), ('d', 35), ('e', 35), ('f', 35)]

*/

import (
	"fmt"
	"sort"
)

// merchant struct: name, maxCapacity, amount
// accounts []merchant
// split(): encapsulate the splitting logic

type (
	merchant struct {
		key         string
		maxCapacity int
		amount      int
	}
)

func split(accounts []*merchant, amount int) []*merchant {
	fmt.Println("accounts", accounts)

	byCapacity := sortByCapacity(accounts) // O(nlogn)

	var result []*merchant

	remainingAmount := amount
	numRemainingMerchants := len(accounts)
	for i := range byCapacity { // O(N)
		splitAmount := remainingAmount / numRemainingMerchants
		if byCapacity[i].maxCapacity > splitAmount {
			// split evenly among remaining merchants
			fillAmount := remainingAmount / numRemainingMerchants
			remainder := remainingAmount % numRemainingMerchants

			byKey := sortByKey(accounts[i:]) // O(nlogn)
			for j := range byKey {
				byKey[j].amount = fillAmount

				if remainder != 0 {
					byKey[j].amount += remainder
					remainder = 0
				}
				result = append(result, byKey[j])
			}
			return result
		}
		byCapacity[i].amount = byCapacity[i].maxCapacity
		result = append(result, byCapacity[i])
		remainingAmount -= byCapacity[i].amount
		numRemainingMerchants--
	}
	return result
}

func sortByKey(accounts []*merchant) []*merchant {
	byKey := make([]*merchant, len(accounts))
	copy(byKey, accounts) // O(N)

	sort.Slice(byKey, func(i, j int) bool { // O(nlogn)
		return byKey[i].key < byKey[j].key
	})
	return byKey
}

func sortByCapacity(accounts []*merchant) []*merchant {
	byCapacity := make([]*merchant, len(accounts))
	copy(byCapacity, accounts) // O(N)

	sort.Slice(byCapacity, func(i, j int) bool { // O(nlogn)
		return byCapacity[i].maxCapacity < byCapacity[j].maxCapacity
	})
	return byCapacity
}
