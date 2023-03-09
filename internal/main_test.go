package main

import "testing"

type TestCase struct {
	Name string
	Step int
}

func Benchmark_Copy(b *testing.B) {

	testCases := []TestCase{
		{Name: "Modify", Step: 1},
		{Name: "Update", Step: 2},
		{Name: "AddItem", Step: 3},
		{Name: "ModifyItem", Step: 4},
	}

	for _, v := range testCases {
		order := NewOrderBuilderWithDefaults().order

		b.Run(v.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProcessOrder_Copy(order, v.Step)
			}
		})
	}
}

func Benchmark_Pointer(b *testing.B) {

	testCases := []TestCase{
		{Name: "Modify", Step: 1},
		{Name: "Update", Step: 2},
		{Name: "AddItem", Step: 3},
		{Name: "ModifyItem", Step: 4},
	}

	for _, v := range testCases {
		order := NewOrderBuilderWithDefaults().order

		b.Run(v.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ProcessOrder_Pointer(&order, v.Step)
			}
		})
	}
}
