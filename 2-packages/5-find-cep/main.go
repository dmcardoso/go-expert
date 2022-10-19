package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	/* Find CEP

	 */

	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer requisição: %v\n", err)
		}

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler resposta: %v\n", err)
		}

		var data CEP
		err = json.Unmarshal(res, &data)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao fazer o parse da resposta: %v\n", err)
		}

		file, err := os.Create("city.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao criar arquivo: %v\n", err)
		}

		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s\nLocalidade: %s\nUF: %s", data.Cep, data.Localidade, data.Uf))

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao escrever em arquivo: %v\n", err)
		}

		fmt.Println("Arquivo criado com sucesso")

		fileContent, err := os.ReadFile("city.txt")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao ler conteúdo do arquivo: %v\n", err)
		}

		fmt.Println(string(fileContent))
	}

}
