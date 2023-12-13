package main

func NegativeFilter(data int) (int, bool) {
	if data < 0 {
		return 0, false
	}
	return data, true
}

func FilterDivisibility3(data int) (int, bool) {
	if data%3 != 0 || data == 0 {
		return 0, false
	}
	return data, true
}
