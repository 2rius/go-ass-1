package main

// Help function Mod helps return the modulo of a number, even if it is negative
func Mod(a, b int) int {
	return (a%b + b) % b
}
