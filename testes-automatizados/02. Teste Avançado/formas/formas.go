package formas

import (
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * (c.raio * c.raio)
}

// func main() {
// 	fmt.Println("Interfaces")

// 	r := retangulo{10, 15}
// 	escreverArea(r)

// 	c := circulo{10}
// 	escreverArea(c)
// }
