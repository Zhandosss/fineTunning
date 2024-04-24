package main

import (
	"fineTunning/providers"
	"fineTunning/questions"
	"fmt"
	"os"
)

const (
	SYSTEM_MSG_TMPL = "In this conversation, you provide information about the API of a provider. The provider is called %s. All information about %[1]s you can get from this site(s): %s"
)

func main() {
	ans, names, sites := providers.Init()
	file, err := os.Create("providerOpenAI.jsonl")
	if err != nil {
		panic("File create:" + err.Error())
	}
	defer file.Close()
	qest := questions.New()

	for i, prov := range ans {
		systemMsg := fmt.Sprintf(SYSTEM_MSG_TMPL, names[i], sites[i])
		for j := range prov {
			questionForProvider := fmt.Sprintf(qest[j], names[i])
			fmt.Fprintf(file, `{"messages":[{"role": "system","content": "%s"}, `, systemMsg)
			fmt.Fprintf(file, `{"role": "user","content": "%s"}, `, questionForProvider)
			fmt.Fprintf(file, `{"role": "assistant","content": "%s"} `, prov[j])
			fmt.Fprint(file, "]}\n")
		}
	}
}
