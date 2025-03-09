package main

import "fmt"

// объявляем интерфейс
type Flyer interface{
	Fly() string
}

func describeFly(object Flyer) string{
	return object.Fly()
}

type Bird struct{
	species string
}

func (b Bird) Fly() string{
	return b.species + " взлетела\n"

}

type Airplane struct{
	model string
}

func (a Airplane) Fly() string{
	return a.model + " взлетел\n"
}

func main() {

	boeing := Airplane{model: "737-800"}
	fmt.Printf(describeFly(boeing))
	pigeon := Bird{species: "pigeon"}
	fmt.Printf(describeFly(pigeon))
}

