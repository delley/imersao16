package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pessoa struct {
	Nome     string
	Idade    int
	Potuacao int
}

type porNome []Pessoa

func (n porNome) Len() int {
	return len(n)
}

func (n porNome) Less(i, j int) bool {
	tki := fmt.Sprintf("%s%d", n[i].Nome, n[i].Idade)
	tkj := fmt.Sprintf("%s%d", n[j].Nome, n[j].Idade)
	return tki < tkj
}

func (n porNome) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {

	args := os.Args[1:]
	fmt.Println(args)
	if len(args) < 2 {
		panic("informe o arquivo de entrada e de saída")
	}
	caminhoOrigem := args[0]
	caminhoDestino := args[1]

	arquivoOrigem, err := os.Open(caminhoOrigem)
	if err != nil {
		panic(err)
	}
	defer arquivoOrigem.Close()

	reader := csv.NewReader(arquivoOrigem)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	pessoas := make([]Pessoa, len(records)-1)
	for m, record := range records {
		if m == 0 {
			continue
		}

		idade, err := strconv.Atoi(strings.TrimSpace(record[1]))
		if err != nil {
			fmt.Println(err)
			continue
		}

		pontuacao, err := strconv.Atoi(strings.TrimSpace(record[2]))
		if err != nil {
			fmt.Println(err)
			continue
		}

		p := Pessoa{
			Nome:     record[0],
			Idade:    idade,
			Potuacao: pontuacao,
		}
		pessoas[m-1] = p
	}

	sort.Sort(porNome(pessoas))

	arquivoDestino, err := os.Create(caminhoDestino)
	if err != nil {
		panic(err)
	}
	defer arquivoDestino.Close()

	writer := csv.NewWriter(arquivoDestino)
	defer writer.Flush()

	headers := []string{"Nome", "Idade", "Pontuação"}

	writer.Write(headers)
	for _, c := range pessoas {
		writer.Write([]string{c.Nome, strconv.Itoa(c.Idade), strconv.Itoa(c.Potuacao)})
	}

	// headers = []string{"carlos", "15", "90"}
	// writer.Write(headers)
	// headers = []string{"Carlos", "22", "75"}
	// writer.Write(headers)
	// headers = []string{"Carlos", "30", "80"}
	// writer.Write(headers)
	// headers = []string{"Joao", "25", "80"}
	// writer.Write(headers)
	// headers = []string{"Maria", "30", "95"}
	// writer.Write(headers)

}
