package main

import "github.com/brianvoe/gofakeit/v6"

type Order struct {
	ID          string
	Name        string
	Description string
	Items       []string

	Field1 string
	Field2 string
	Field3 string
}

type OrderBuilder struct {
	order Order
}

func NewOrderBuilderWithDefaults() OrderBuilder {
	order := Order{
		ID:          gofakeit.UUID(),
		Name:        gofakeit.Name(),
		Description: gofakeit.Paragraph(1, 5, 10, " "),
		Items:       BuildArray(100),
		Field1:      gofakeit.Word(),
		Field2:      gofakeit.Word(),
		Field3:      gofakeit.Word(),
	}

	return OrderBuilder{
		order: order,
	}
}

type User struct {
}
