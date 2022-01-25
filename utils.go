package main

func minMaxDefault64(value int64, min int64, max int64) int64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}

	return value
}

func minMaxDefault(value int, min int, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}

	return value
}
