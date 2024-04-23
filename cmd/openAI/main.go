package main

import (
	"fineTunning/providers"
	"fineTunning/questions"
	"fmt"
	"os"
)

func main() {
	ans, names := providers.Init()
	file, err := os.Create("providerOpenAI.jsonl")
	if err != nil {
		panic("File create" + err.Error())
	}
	defer file.Close()
	qest := questions.New()

	for i, prov := range ans {
		fmt.Fprintf(file, `{"messages":[{"role": "system","text": "In this conversation, you provide information about the API of a provider. The provider is called %s"},`, names[i])
		for j := range prov {
			questionForProvider := fmt.Sprintf(qest[j], names[i])
			fmt.Fprintf(file, `{"role": "user","text": "%s"},`, questionForProvider)
			fmt.Fprintf(file, `{"role": "assistant","text": "%s"}`, prov[j])
			if j != len(prov)-1 {
				fmt.Fprint(file, `,`)
			}
		}
		fmt.Fprint(file, "]}\n")
	}

	for i := 0; i < 10-len(names); i++ {
		fmt.Fprintf(file, `{"messages":[{"role": "user","text": "bye"},{"role": "assistant","text": "bye"}]}`)
		fmt.Fprintln(file)

	}
}
