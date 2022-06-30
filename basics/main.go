package main //nombre del paquete
//package cualquier nombre del paquete

import "fmt" 


func main() {
	var mensaje string = "Hola Mundo"
	mensajeFacil := "Hola Mundo usando :="
	fmt.Println(mensaje)
	fmt.Println(mensajeFacil)

	a :=10.
	var b float64 = 3
	fmt.Println(a/b)

	var c int = 10
	d := 3
	fmt.Println(c/d)

	x:= true
	y:= true

	fmt.Println(x||y)
	fmt.Println(x&&y)
	fmt.Println(!x)

	privada()

	Publica()

	imprimirHola()

}

func privada()  {
	fmt.Println("Ejecutar logica de negocio que debe ser exportada")
}
func Publica()  {
	fmt.Println("Lógica que quiero exportar")
}
func imprimirHola()  {
	defer fmt.Println("Mundo!")
	fmt.Print("Hola ")
}
