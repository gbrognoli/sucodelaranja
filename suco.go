package main

import (
	"fmt"
)

type Juice struct {
	Name     string
	Quantity float64
}

var juiceMap map[string]Juice

func main() {
	juiceMap = make(map[string]Juice)

	// Adicionar suco
	addJuice("Laranja", 2.5)
	addJuice("Limão", 1.8)

	// Atualizar suco
	updateJuice("Laranja", 3.2)

	// Remover suco
	removeJuice("Limão")

	// Exibir sucos
	displayJuices()
}

func addJuice(name string, quantity float64) {
	juice := Juice{
		Name:     name,
		Quantity: quantity,
	}
	juiceMap[name] = juice
	fmt.Printf("Suco %s adicionado com sucesso!\n", name)
}

func updateJuice(name string, quantity float64) {
	if _, ok := juiceMap[name]; ok {
		juiceMap[name].Quantity = quantity
		fmt.Printf("Suco %s atualizado com sucesso!\n", name)
	} else {
		fmt.Printf("Suco %s não encontrado!\n", name)
	}
}

func removeJuice(name string) {
	if _, ok := juiceMap[name]; ok {
		delete(juiceMap, name)
		fmt.Printf("Suco %s removido com sucesso!\n", name)
	} else {
		fmt.Printf("Suco %s não encontrado!\n", name)
	}
}

func displayJuices() {
	fmt.Println("Lista de sucos:")
	for name, juice := range juiceMap {
		fmt.Printf("- %s: %.2f litros\n", name, juice.Quantity)
	}
}
