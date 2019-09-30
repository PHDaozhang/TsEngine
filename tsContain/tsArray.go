package tsContain

func InArrayInt(array []int64, key int64) bool {
	for _, v := range array {
		if key == v {
			return true
		}
	}

	return false
}

func InArrayString(array []string, key string) bool {
	for _, v := range array {
		if key == v {
			return true
		}
	}

	return false
}

func InArrayIntN(array []int, key int) bool {
	for _, v := range array {
		if key == v {
			return true
		}
	}

	return false
}
