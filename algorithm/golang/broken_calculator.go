package main

func brokenCalc(X int, Y int) int {
	if X >= Y {
		return X - Y
	}
	if Y % 2 == 0 {
		return brokenCalc(X, Y / 2) + 1
	}
	return brokenCalc(X, Y + 1) + 1
}
