package main

import (
	"fineTunning/providers"
	"fineTunning/questions"
	"fmt"
	"os"
)

func main() {
	ans, names := providers.Init()
	file, err := os.Create("providerYandex.jsonl")
	if err != nil {
		panic("File create" + err.Error())
	}
	defer file.Close()
	qest := questions.New()
	for i, prov := range ans {
		for j := range prov {
			questionForProvider := fmt.Sprintf(qest[j], names[i])
			fmt.Fprintf(file, `{"request": [{"role": "system", "text": "In this conversation, you provide information about the API of a provider. The provider is called %s"},{"role": "user", "text": "%s"}], "response": "%s"}`+"\n", names[i], questionForProvider, prov[j])
		}
	}
}
