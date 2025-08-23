package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	//ler o arquivo csv
	//printar as perguntas na tela
	//receber a resposta do usuario
	//verificar se a resposta do usuario e igual a do arquivo
	//salvar se esta certo ou errada
	//printar quantidade correta e incorreta

	fmt.Println("quiz game start")

	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close() //defer roda no final da execucao da funcao

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(records)
}
