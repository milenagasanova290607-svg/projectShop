package main

import "fmt"

type Product struct {
	ID       int
	Name     string
	Category string
	Price    int
	Size     string
	Quantity int
}

func main() {
	product1 := Product{
		ID:       1,
		Name:     "Футболка белая",
		Category: "Футболки",
		Price:    600,
		Size:     "M",
		Quantity: 2000,
	}

	product2 := Product{
		ID:       2,
		Name:     "Джинсы серые",
		Category: "Джинсы",
		Price:    2000,
		Size:     "M",
		Quantity: 2500,
	}
	fmt.Println("Товар №1")
	fmt.Println("ID товара:", product1.ID)
	fmt.Println("Название товара:", product1.Name)
	fmt.Println("Категория товара:", product1.Category)
	fmt.Println("Цена:", product1.Price)
	fmt.Println("Размер:", product1.Size)
	fmt.Println("Количество:", product1.Quantity)

	fmt.Println("Товар №2")
	fmt.Println("ID товара:", product2.ID)
	fmt.Println("Название товара:", product2.Name)
	fmt.Println("Категория товара:", product2.Category)
	fmt.Println("Цена:", product2.Price)
	fmt.Println("Размер:", product2.Size)
	fmt.Println("Количество:", product2.Quantity)

	product1.Price = 700.00

	fmt.Println("Новая цена футболки:", product1.Price)
}
