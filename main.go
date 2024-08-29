package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const tamanhoMaxManual = 10

// Função para gerar uma matriz aleatória
func geraMatriz(linhas, colunas int) [][]int {
	matriz := make([][]int, linhas)
	for i := range matriz {
		matriz[i] = make([]int, colunas)
		for j := range matriz[i] {
			matriz[i][j] = rand.Intn(100) // Valores aleatórios entre 0 e 99
		}
	}
	return matriz
}

// Função para exibir uma matriz
func imprimeMatriz(matriz [][]int) {
	for _, linha := range matriz {
		fmt.Println(linha)
	}
}

// Função para somar duas matrizes sequencialmente
func somaMatrizesSequencial(matrizA, matrizB [][]int) [][]int {
	linhas := len(matrizA)
	colunas := len(matrizA[0])
	resultado := make([][]int, linhas)
	for i := 0; i < linhas; i++ {
		resultado[i] = make([]int, colunas)
		for j := 0; j < colunas; j++ {
			resultado[i][j] = matrizA[i][j] + matrizB[i][j]
		}
	}
	return resultado
}

// Função para somar duas matrizes em paralelo
func somaMatrizesParalelo(matrizA, matrizB [][]int) [][]int {
	linhas := len(matrizA)
	colunas := len(matrizA[0])
	resultado := make([][]int, linhas)
	var wg sync.WaitGroup

	for i := 0; i < linhas; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resultado[i] = make([]int, colunas)
			for j := 0; j < colunas; j++ {
				resultado[i][j] = matrizA[i][j] + matrizB[i][j]
			}
		}(i)
	}

	wg.Wait()
	return resultado
}

// Função para multiplicar duas matrizes sequencialmente
func multiplicaMatrizesSequencial(matrizA, matrizB [][]int) [][]int {
	linhas := len(matrizA)
	colunas := len(matrizB[0])
	comum := len(matrizB)
	resultado := make([][]int, linhas)
	for i := 0; i < linhas; i++ {
		resultado[i] = make([]int, colunas)
		for j := 0; j < colunas; j++ {
			soma := 0
			for k := 0; k < comum; k++ {
				soma += matrizA[i][k] * matrizB[k][j]
			}
			resultado[i][j] = soma
		}
	}
	return resultado
}

// Função para multiplicar duas matrizes em paralelo
func multiplicaMatrizesParalelo(matrizA, matrizB [][]int) [][]int {
	linhas := len(matrizA)
	colunas := len(matrizB[0])
	comum := len(matrizB)
	resultado := make([][]int, linhas)
	var wg sync.WaitGroup

	for i := 0; i < linhas; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resultado[i] = make([]int, colunas)
			for j := 0; j < colunas; j++ {
				soma := 0
				for k := 0; k < comum; k++ {
					soma += matrizA[i][k] * matrizB[k][j]
				}
				resultado[i][j] = soma
			}
		}(i)
	}

	wg.Wait()
	return resultado
}

// Função para transpor uma matriz sequencialmente
func transpoeMatrizSequencial(matriz [][]int) [][]int {
	linhas := len(matriz)
	colunas := len(matriz[0])
	resultado := make([][]int, colunas)
	for i := 0; i < colunas; i++ {
		resultado[i] = make([]int, linhas)
		for j := 0; j < linhas; j++ {
			resultado[i][j] = matriz[j][i]
		}
	}
	return resultado
}

// Função para transpor uma matriz em paralelo
func transpoeMatrizParalelo(matriz [][]int) [][]int {
	linhas := len(matriz)
	colunas := len(matriz[0])
	resultado := make([][]int, colunas)
	var wg sync.WaitGroup

	for i := 0; i < colunas; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resultado[i] = make([]int, linhas)
			for j := 0; j < linhas; j++ {
				resultado[i][j] = matriz[j][i]
			}
		}(i)
	}

	wg.Wait()
	return resultado
}

func main() {
	var escolha int
	var linhasA, colunasA, linhasB, colunasB int

	fmt.Println("Escolha a operação que deseja realizar:")
	fmt.Println("1 - Soma de Matrizes")
	fmt.Println("2 - Multiplicação de Matrizes")
	fmt.Println("3 - Transposição de Matriz")
	fmt.Print("Digite o número da operação: ")
	fmt.Scan(&escolha)

	fmt.Println("Deseja inserir as dimensões das matrizes manualmente ou gerar aleatoriamente?")
	fmt.Println("1 - Inserir dimensões manualmente")
	fmt.Println("2 - Gerar aleatoriamente")
	fmt.Print("Escolha: ")
	var escolhaInput int
	fmt.Scan(&escolhaInput)

	if escolhaInput == 1 {
		fmt.Println("Digite o número de linhas para a matriz A (máximo 10): ")
		fmt.Scan(&linhasA)
		if linhasA < 1 || linhasA > tamanhoMaxManual {
			fmt.Println("Número de linhas inválido.")
			return
		}

		switch escolha {
		case 1, 3:
			fmt.Println("Digite o número de colunas para as matrizes (máximo 10): ")
			fmt.Scan(&colunasA)
			if colunasA < 1 || colunasA > tamanhoMaxManual {
				fmt.Println("Número de colunas inválido.")
				return
			}
			linhasB, colunasB = linhasA, colunasA // Para soma e transposição, dimensões são as mesmas
		case 2:
			fmt.Println("Digite o número de colunas para Matriz A (isso será igual ao número de linhas de Matriz B) (máximo 10): ")
			fmt.Scan(&colunasA)
			if colunasA < 1 || colunasA > tamanhoMaxManual {
				fmt.Println("Número de colunas inválido.")
				return
			}

			linhasB = colunasA // O número de linhas de B deve ser igual ao número de colunas de A

			fmt.Println("Digite o número de colunas para Matriz B (máximo 10): ")
			fmt.Scan(&colunasB)
			if colunasB < 1 || colunasB > tamanhoMaxManual {
				fmt.Println("Número de colunas inválido.")
				return
			}
		}
	} else {
		fmt.Println("Escolha o tamanho da matriz gerada:")
		fmt.Println("1 - 100x100")
		fmt.Println("2 - 1000x1000")
		fmt.Println("3 - 10000x10000")
		fmt.Println("4 - 100000x100000")
		fmt.Print("Escolha: ")
		var escolhaTamanho int
		fmt.Scan(&escolhaTamanho)

		switch escolhaTamanho {
		case 1:
			linhasA, colunasA, linhasB, colunasB = 100, 100, 100, 100
		case 2:
			linhasA, colunasA, linhasB, colunasB = 1000, 1000, 1000, 1000
		case 3:
			linhasA, colunasA, linhasB, colunasB = 10000, 10000, 10000, 10000
		case 4:
			linhasA, colunasA, linhasB, colunasB = 100000, 100000, 100000, 100000
		default:
			fmt.Println("Escolha inválida.")
			return
		}
	}

	var matrizA, matrizB [][]int

	// Gerar as matrizes com base nas dimensões fornecidas
	matrizA = geraMatriz(linhasA, colunasA)
	if escolha == 1 || escolha == 2 { // Soma ou Multiplicação
		matrizB = geraMatriz(linhasB, colunasB)
	}

	if escolhaInput == 1 {
		fmt.Println("Matriz A:")
		imprimeMatriz(matrizA)
		if escolha == 1 || escolha == 2 {
			fmt.Println("Matriz B:")
			imprimeMatriz(matrizB)
		}
	}

	// Executa e mede o tempo da operação sequencial
	inicioSequencial := time.Now()
	var resultadoSequencial [][]int

	switch escolha {
	case 1:
		fmt.Println("Realizando a soma de matrizes sequencialmente...")
		resultadoSequencial = somaMatrizesSequencial(matrizA, matrizB)
	case 2:
		fmt.Println("Realizando a multiplicação de matrizes sequencialmente...")
		resultadoSequencial = multiplicaMatrizesSequencial(matrizA, matrizB)
	case 3:
		fmt.Println("Realizando a transposição de matriz sequencialmente...")
		resultadoSequencial = transpoeMatrizSequencial(matrizA)
	}
	tempoSequencial := time.Since(inicioSequencial)
	fmt.Printf("Tempo de execução sequencial: %s\n", tempoSequencial)

	// Executa e mede o tempo da operação paralela
	inicioParalelo := time.Now()
	var resultadoParalelo [][]int

	switch escolha {
	case 1:
		fmt.Println("Realizando a soma de matrizes em paralelo...")
		resultadoParalelo = somaMatrizesParalelo(matrizA, matrizB)
	case 2:
		fmt.Println("Realizando a multiplicação de matrizes em paralelo...")
		resultadoParalelo = multiplicaMatrizesParalelo(matrizA, matrizB)
	case 3:
		fmt.Println("Realizando a transposição de matriz em paralelo...")
		resultadoParalelo = transpoeMatrizParalelo(matrizA)
	}
	tempoParalelo := time.Since(inicioParalelo)
	fmt.Printf("Tempo de execução paralelo: %s\n", tempoParalelo)

	if escolhaInput == 1 {
		switch escolha {
		case 1:
			fmt.Println("Resultado da soma:")
		case 2:
			fmt.Println("Resultado da multiplicação:")
		case 3:
			fmt.Println("Resultado da transposição:")
		}
		imprimeMatriz(resultadoParalelo)
	}

	_ = resultadoSequencial
	_ = resultadoParalelo
}
