package main

import "fmt"

func calculatePrice(price int, quantity int) int {
	return price * quantity
}

func main() {

	shopName := "Онлайн магазин одежды"
	productName := "Женская белая футболка"
	price := 600
	quantity := 3000
	buyQuantity := 6
	if buyQuantity <= quantity {
		fmt.Println("Заказ успешно оформлен")
	} else {
		fmt.Print("Недостаточно товара на складе")
		return
	}
	totalPrice := calculatePrice(price, buyQuantity)
	fmt.Println("Стоимость заказа:", totalPrice)
	remainingQuantity := quantity - buyQuantity

	fmt.Println("Название магазина:", shopName)
	fmt.Println("Название товара:", productName)
	fmt.Println("Цена товара:", price)
	fmt.Println("Количество товара:", quantity)
	fmt.Println("Количество заказа покупателя:", buyQuantity)
	fmt.Println("Стоимость заказа:", totalPrice)
	fmt.Println("Остаток на складе:", remainingQuantity)
	for i := 1; i <= buyQuantity; i++ {
		fmt.Println("Футболка №", i)
	}

	products := []string{"Футболка", "Джинсы", "Лонгслив", "Юбка", "Толстовка", "Зип-худи", "Сапоги", "Кроссовки"}
	products = append(products, "Аксессуары")
	for m := 0; m < len(products); m++ {
		fmt.Println(m+1, products[m])
	}
	fmt.Println("Всего товаров:", len(products))

	prices := make(map[string]int)
	prices["Футболка"] = 600
	prices["Джинсы"] = 2000
	prices["Лонгслив"] = 1000
	prices["Юбка"] = 1000
	prices["Толстовка"] = 1500
	prices["Зип-худи"] = 1500
	prices["Сапоги"] = 4000
	prices["Кроссовки"] = 3000
	prices["Аксессуары"] = 300

	if v, done := prices["Джинсы"]; done {
		fmt.Println("Цена джинсов:", v)
	} else {
		fmt.Println("Товара нет")
	}
	fmt.Println("Каталог товаров:")
	for t, p := range prices {
		fmt.Println("Товар:", t, "Цена товара:", p)
	}

	fmt.Println("Всего товаров:", len(prices))
}
