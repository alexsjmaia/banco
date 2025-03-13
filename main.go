package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque relizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Printf("%.2f\n", contaDaSilvia.saldo)

	valorDoSaque := 499.99
	fmt.Println(contaDaSilvia.Sacar(float64(valorDoSaque)))

	fmt.Printf("Saldo Atual R$ %.2f\n", contaDaSilvia.saldo)

}
