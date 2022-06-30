package main

import "fmt"

type perro struct {
}

type pez struct {
}

type pajaro struct {
}

func (perro) caminar() string {
	return "soy un perro y camino"
}

func (pez) nadar() string {
	return "soy un pez y nado"
}

func (pajaro) volar() string {
	return "soy un pajaro y vuelo"
}

// aqui se ve un codigo repetitivo, sin interfaz

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(p pez) {
	fmt.Println(p.nadar())
}

func moverPajaro(p pajaro){
	fmt.Println(p.volar())
}

func main() {
	dog := perro{}

	fish := pez{}

	bird := pajaro{}

	//
	moverPerro(dog)
	moverPez(fish)
	moverPajaro(bird)

}
