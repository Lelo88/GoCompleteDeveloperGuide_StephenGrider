package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d:= newDeck() // creamos un nuevo mazo. Esperamos 16 cartas porque tenemos cuatro tipos de palos y 
					// cuatro valores por cada uno

	// comprobamos que el mazo tiene 16 cartas
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d)) // cadena formateada
	}

	// comprobamos que la primer carta del mazo sea un as de espadas
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	// comprobamos que la última carta sea el 4 de trébol
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting") // borramos el archivo si existe

	deck := newDeck() // creamos un nuevo mazo
	deck.saveToFile("_decktesting") // guardamos el mazo en el archivo

	loadedDeck := newDeckFromFile("_decktesting") // cargamos el mazo desde el archivo

	if len(loadedDeck) != 16 { // podemos cambiar el valor de 16 para ver si se produce el error
		t.Errorf("Expected 16 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}