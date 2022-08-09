package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoJoao := ContaCorrente{titular: "João", numeroAgencia: 0001, numeroConta: 123456, saldo: 10000.0}
	contaDaMonica := ContaCorrente{titular: "Monica", numeroAgencia: 221, numeroConta: 56789, saldo: 20000.0}

	fmt.Println(contaDoJoao)
	fmt.Println(contaDaMonica)
	
}