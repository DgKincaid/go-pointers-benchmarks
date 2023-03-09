package main

import (
	"github.com/brianvoe/gofakeit/v6"
)

func BuildArray(l int) []string {
	arr := make([]string, l)

	for i := 0; i < l; i++ {
		arr[i] = gofakeit.Word()
	}

	return arr
}
