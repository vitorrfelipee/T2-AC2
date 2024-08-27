package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func geraMatriz(linhas, colunas int) [][]int {
	matriz := make([][]int, linhas)
	for i := range matriz {
		matriz[i] = make([]int, colunas)
		for j := range matriz[i] {
			matriz[i][j] = rand.Intn(100)
		}
	}
	return matriz
}

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

func medirTempoExecucao(operacao func() [][]int) time.Duration {
	inicio := time.Now()
	_ = operacao()
	return time.Since(inicio)
}

func main() {
	tamanhos := []int{100, 1000}
	repeticoes := 10

	arquivo, err := os.Create("tempos_execucao.csv")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer arquivo.Close()

	writer := csv.NewWriter(arquivo)
	defer writer.Flush()

	// Escrevendo o cabeçalho no CSV
	writer.Write([]string{"Tamanho", "Operacao", "Tipo", "Tempo_medio"})

	for _, tamanho := range tamanhos {
		for i := 0; i < repeticoes; i++ {
			matrizA := geraMatriz(tamanho, tamanho)
			matrizB := geraMatriz(tamanho, tamanho)

			// Soma Sequencial
			tempoSeqSoma := medirTempoExecucao(func() [][]int {
				return somaMatrizesSequencial(matrizA, matrizB)
			})
			writer.Write([]string{strconv.Itoa(tamanho), "Soma", "Sequencial", strconv.FormatFloat(tempoSeqSoma.Seconds(), 'f', 6, 64)})

			// Soma Paralela
			tempoParSoma := medirTempoExecucao(func() [][]int {
				return somaMatrizesParalelo(matrizA, matrizB)
			})
			writer.Write([]string{strconv.Itoa(tamanho), "Soma", "Paralelo", strconv.FormatFloat(tempoParSoma.Seconds(), 'f', 6, 64)})

			// Multiplicação Sequencial
			tempoSeqMult := medirTempoExecucao(func() [][]int {
				return multiplicaMatrizesSequencial(matrizA, matrizB)
			})
			writer.Write([]string{strconv.Itoa(tamanho), "Multiplicacao", "Sequencial", strconv.FormatFloat(tempoSeqMult.Seconds(), 'f', 6, 64)})

			// Multiplicação Paralela
			tempoParMult := medirTempoExecucao(func() [][]int {
				return multiplicaMatrizesParalelo(matrizA, matrizB)
			})
			writer.Write([]string{strconv.Itoa(tamanho), "Multiplicacao", "Paralelo", strconv.FormatFloat(tempoParMult.Seconds(), 'f', 6, 64)})

			// Transposição Sequencial
			tempoSeqTrans := medirTempoExecucao(func() [][]int {
				return transpoeMatrizSequencial(matrizA)
			})
			writer.Write([]string{strconv.Itoa(tamanho), "Transposicao", "Sequencial", strconv.FormatFloat(tempoSeqTrans.Seconds(), 'f', 6, 64)})

			// Transposição Paralela
			tempoParTrans := medirTempoExecucao(func() [][]int {
				return transpoeMatrizParalelo(matrizA)
			})
			writer.Write([]string{strconv.Itoa(tamanho), "Transposicao", "Paralelo", strconv.FormatFloat(tempoParTrans.Seconds(), 'f', 6, 64)})
		}
	}

	fmt.Println("Os tempos de execução foram salvos no arquivo tempos_execucao.csv.")
}
