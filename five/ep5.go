package five

func CalculateSmallestMultipleOfOneToTwenty() int {
	smallestValue := 2520
	var valuesToCheck = [...]int{11, 13, 14, 16, 17, 18, 19, 20}
	for true {
		cleanlyDivides := true
		for _, i := range valuesToCheck {
			if smallestValue%i != 0 {
				cleanlyDivides = false
				break
			}
		}

		if cleanlyDivides {
			break
		}

		smallestValue++
	}
	return smallestValue
}
