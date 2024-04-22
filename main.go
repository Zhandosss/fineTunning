package main

import (
	"fineTunning/questions"
	"fineTunning/training"
	"fmt"
	"os"
)

type provider interface {
	GetAns() []string
	GetName() string
	GenerateAns()
}

func initProviders() ([][]string, []string) {
	providers := make([][]string, 0, 10)
	names := make([]string, 0, 10)
	mural := training.NewMural()
	mural.GenerateAns()
	names = append(names, mural.GetName())
	providers = append(providers, mural.GetAns())
	wrike := training.NewWrike()
	wrike.GenerateAns()
	names = append(names, wrike.GetName())
	providers = append(providers, wrike.GetAns())
	return providers, names
}

func main() {
	providers, names := initProviders()
	file, err := os.Create("provider.jsonl")
	if err != nil {
		panic("File create" + err.Error())
	}
	defer file.Close()
	fmt.Println(len(providers[0]))
	fmt.Println(len(names))
	qest := questions.New()

	for i, prov := range providers {
		fmt.Fprintf(file, `{"messages":[{"role": "system","content": "In this conversation, you provide information about the API of a provider. The provider is called %s"},`, names[i])
		for j := range prov {
			fmt.Fprintf(file, `{"role": "user","content": "%s"},`, qest[j])
			fmt.Fprintf(file, `{"role": "assistant","content": "%s"}`, prov[j])
			if j != len(prov)-1 {
				fmt.Fprint(file, `,`)
			}
		}
		fmt.Fprint(file, "]}\n")
	}
}
