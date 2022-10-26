package main

import (
	"fmt"
	"math/rand"
	"time"
)

func StringGenerator(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func main() {
	for {
		fmt.Println(StringGenerator(5))
		time.Sleep(time.Second)
	}
}
