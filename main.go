package main

import "fmt"

func main() {
	fmt.Println("main")
}

func ProcessOrder_Copy(order Order, step int) Order {
	var updatedOrder Order

	switch step {
	case 1:
		updatedOrder = ModifyStep_Copy(order)
	case 2:
		updatedOrder = UpdateStep_Copy(order)
	case 3:
		updatedOrder = AddItem_Copy(order)
	case 4:
		updatedOrder = ModifyItem_Copy(order)
	}

	return updatedOrder
}

func ModifyStep_Copy(order Order) Order {
	order.Description += " updated"
	return order
}

func UpdateStep_Copy(order Order) Order {
	order.Description = "override"
	return order
}

func AddItem_Copy(order Order) Order {
	order.Items = append(order.Items, "new item")
	return order
}

func ModifyItem_Copy(order Order) Order {
	order.Items[0] = "new values"
	return order
}

func ProcessOrder_Pointer(order *Order, step int) *Order {
	var updatedOrder *Order

	switch step {
	case 1:
		updatedOrder = ModifyStep_Pointer(order)
	case 2:
		updatedOrder = UpdateStep_Pointer(order)
	case 3:
		updatedOrder = AddItem_Pointer(order)
	case 4:
		updatedOrder = ModifyItem_Pointer(order)
	}
	return updatedOrder
}

func ModifyStep_Pointer(order *Order) *Order {
	order.Description += " updated"
	return order
}

func UpdateStep_Pointer(order *Order) *Order {
	order.Description = "override"
	return order
}

func AddItem_Pointer(order *Order) *Order {
	order.Items = append(order.Items, "new item")
	return order
}

func ModifyItem_Pointer(order *Order) *Order {
	order.Items[0] = "new values"
	return order
}

func NoReturn_Modify_Pointer(order *Order) {
	order.Description += " updated"
}
