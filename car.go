package main

import (
	"errors"
	"fmt"
	// "os"
)

type Car struct {
	model     string
	speed     float32
	max_speed int
}

func (c *Car) Razgon(razgon float32) error {
	if razgon > float32(c.max_speed) {
		return errors.New("Превышение максимальной скорости!")
	}

	c.speed += razgon
	switch {
	case c.speed <= 60:
		fmt.Printf("Вы разогнались до %.2f км/ч. Аккуратно!\n", c.speed)

	case c.speed >= 60:
		fmt.Printf("Вы разогнались до %.2f км/ч. Вы превышаете!\n", c.speed)
	}
	return nil

}

func (c *Car) Brake(stop float32) {
	c.speed -= stop

	if c.speed <= 0 {
		c.speed = 0
	}

	switch {
	case c.speed > 0:
		fmt.Printf("Вы замедлились до %.2f км/ч.\n", c.speed)

	case c.speed <= 0:
		fmt.Printf("Мы остановились\n")
	}
}

func main() {
	car := Car{model: "Toyota", max_speed: 210}
	fmt.Printf("Вы стартовали. Текущая скорость: %.2f км/ч.\n", car.speed)

	err := car.Razgon(20)
	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("Текущая скорость: %.2f.\n", car.speed)
	}

	car.Brake(10)

}
