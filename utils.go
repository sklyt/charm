package main

func max[T ~uint32](a, b T) T {
	if a > b {
		return a
	}
	return b
}
