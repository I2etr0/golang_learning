package main

import (
	"fmt"
)

type House struct {
	color_wall string
	color_roof string
}

func change_wall_color(colour string, h *House) {
	h.color_wall = colour
	fmt.Println("Цвет стен после покраски внутри функции:", h.color_wall)
}

func main() {
	h := House{color_wall: "White"}
	fmt.Println("Цвет стен до покраски:", h.color_wall)
	change_wall_color("Red", &h)
	fmt.Println("Цвет стен после покраски:", h.color_wall)
}
