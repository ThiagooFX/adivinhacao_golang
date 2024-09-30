package main

import (
	"fmt"
	"math/rand"
	"time"
)

func reiniciar() bool {
	fmt.Println("Deseja jogar novamente? (s/n)")
	var resposta string
	fmt.Scan(&resposta)
	if resposta == "s" {
		return true
	} else if resposta == "n" {
		fmt.Println("Obrigado por jogar!")
		return false
	} else {
		fmt.Println("Opção inválida")
		return reiniciar()
	}
}

func jogo() {
	rand.Seed(time.Now().UnixNano()) // Define a semente para gerar números aleatórios
	numerosecreto := rand.Intn(100) + 1 // Gera um número entre 1 e 100

	var limitetentativa int
	var dificuldade int
	var tentativa int

	fmt.Println("Bem-vindo ao jogo de adivinhação!")
	fmt.Println("Escolha o nível de dificuldade:")
	fmt.Println("1 = Fácil - 20 tentativas")
	fmt.Println("2 = Médio - 10 tentativas")
	fmt.Println("3 = Difícil - 5 tentativas")
	fmt.Scan(&dificuldade)

	switch dificuldade {
	case 1:
		limitetentativa = 20
	case 2:
		limitetentativa = 10
	case 3:
		limitetentativa = 5
	default:
		fmt.Println("Opção inválida")
		return
	}

	fmt.Println("Escolha um número entre 1 e 100")

	for i := 0; i < limitetentativa; i++ {
		fmt.Printf("Tentativa %d de %d \n", i+1, limitetentativa)
		fmt.Scan(&tentativa) // Ler a tentativa do usuário

		if tentativa < numerosecreto {
			fmt.Println("O número secreto é maior.")
		} else if tentativa > numerosecreto {
			fmt.Println("O número secreto é menor.")
		} else {
			fmt.Println("Parabéns! Você acertou o número!")
			return
		}
	}

	fmt.Printf("Você atingiu o limite de tentativas. O número era %d. \n", numerosecreto)
}

func main() {
	for {
		jogo() // Executa o jogo principal
		if !reiniciar() {
			break
		}
	}
}