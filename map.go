package main

import "fmt"

func main() {

	// иницаилизания меню
	// меню как в python словарь ключ:значение
	menu := map[string]float32{
		"Halal": 150,
		"Hleb":  12.50,
		"Voda":  0.0,
	}

	fmt.Println(menu)
}
