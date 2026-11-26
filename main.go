package main

import "fmt"

type Money = int
type ID = int
type Role = string

type Product struct {
	ID       ID
	Name     string
	Category string
	Price    Money
	Size     string
	Quantity int
}

type User struct {
	ID    int
	Name  string
	Login string
	Role  Role
}

type Informer interface {
	showInfo()
}

func printInfo(inf Informer) {
	inf.showInfo()
}

func (p Product) showInfo() {
	fmt.Println("ID товара:", p.ID)
	fmt.Println("Название товара:", p.Name)
	fmt.Println("Категория товара:", p.Category)
	fmt.Println("Цена:", p.Price)
	fmt.Println("Размер:", p.Size)
	fmt.Println("Количество:", p.Quantity)
}

func (u User) showInfo() {
	fmt.Println("ID покупателя:", u.ID)
	fmt.Println("Имя покупателя:", u.Name)
	fmt.Println("Логин:", u.Login)
	fmt.Println("Роль:", u.Role)
}

func (p *Product) changePrice(newPrice int) {
	p.Price = newPrice
}

func (p *Product) changeQuantity(newQuantity int) {
	p.Quantity = newQuantity
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

	user1 := User{
		ID:    1,
		Name:  "Милена",
		Login: "milena2906",
		Role:  "client",
	}

	user2 := User{
		ID:    2,
		Name:  "Александра",
		Login: "nalazet",
		Role:  "client",
	}

	printInfo(product1)

	printInfo(product2)

	printInfo(user1)

	printInfo(user2)

	product1.changePrice(800)

	product2.changeQuantity(1500)

	fmt.Println("Новая цена футболки:", product1.Price)

	fmt.Println("Новое количество джинсов:", product2.Quantity)

}
