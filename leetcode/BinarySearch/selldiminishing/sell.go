package selldiminishing

func maxProfit(inventory []int, orders int) int {
	maxItem, thresholdValue, count, res, mod := 0, -1, 0, 0, 1000000007

	for i := 0; i < len(inventory); i++ {
		if inventory[i] > maxItem {
			maxItem = inventory[i]
		}
	}

	low, high := 0, maxItem
	for low <= high {
		mid := low + ((high - low) >> 1)
		for i := 0; i < len(inventory); i++ {
			count += max(inventory[i]-mid, 0)
		}

		if count <= orders {
			thresholdValue = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
		count = 0
	}
	count = 0
	for i := 0; i < len(inventory); i++ {
		count += max(inventory[i]-thresholdValue, 0)
	}

	count = orders - count

	for i := 0; i < len(inventory); i++ {
		if inventory[i] >= thresholdValue {
			if count > 0 {
				res += (thresholdValue + inventory[i]) * (inventory[i] - thresholdValue + 1) / 2
				count--
			} else {
				res += (thresholdValue + 1 + inventory[i]) * (inventory[i] - thresholdValue) / 2
			}
		}
	}

	return res % mod
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
