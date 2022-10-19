package main

import (
	"encoding/json"
	"io"
	"net/http"
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
	http.HandleFunc("/", FindCEPHandler)
	http.ListenAndServe(":8080", nil)
}

func FindCEPHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		responseWriter.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := request.URL.Query().Get("cep")
	if cepParam == "" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	cep, err := SearchCep(cepParam)

	if err != nil {
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusOK)

	// result, err := json.Marshal(cep)
	// if err != nil {
	// 	responseWriter.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }
	// responseWriter.Write(result)

	// ^ Equals to ->

	json.NewEncoder(responseWriter).Encode(cep)
}

func SearchCep(cep string) (*CEP, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	var cepData CEP
	err = json.Unmarshal(body, &cepData)

	if err != nil {
		return nil, err
	}

	return &cepData, nil
}
