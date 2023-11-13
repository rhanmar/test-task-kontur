package main

import "math/rand"

// GetRandomIntForRange - получить случайное целое число в указанных пределах
func GetRandomIntForRange(max int, min int) int {
	return rand.Intn(max-min) + min
}
