package main

// import "fmt"

func main() {
	//si coloco el nombre equivocado al abrir el archivo, me arrojará el error que hayamos indicado
	//cards := newDeckFromFile("mynewdeck")
	//cards.print()
	cards := newDeck() // creo un nuevo mazo
	cards.shuffle() // mezclo el mazo
	cards.print() // se imprime el nuevo mazo mezclado por la función shuffle
} 