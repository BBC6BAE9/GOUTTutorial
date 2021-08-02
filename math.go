package main

func MyMod(a int, b int) int{
	if b == 0 {
		return -1
	}
	return a % b
}
