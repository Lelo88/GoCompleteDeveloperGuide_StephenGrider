package main

// import "fmt"

func main() {
	//si coloco el nombre equivocado al abrir el archivo, me arrojará el error que hayamos indicado
	cards := newDeckFromFile("mynewdeck")
	cards.print()
}