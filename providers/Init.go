package providers

import (
	"fineTunning/training"
	"strings"
)

type Iprovider interface {
	GenerateAns()
	GetName() string
	GetAns() []string
	GetSearchSites() []string
}

func Init() ([][]string, []string, []string) {
	ans := make([][]string, 0, 5)
	names := make([]string, 0, 5)
	sites := make([]string, 0, 5)

	providers := make([]Iprovider, 0, 5)
	providers = append(providers, training.NewMural(), training.NewWrike(), training.NewKlaviyo(), training.NewRingCentral(), training.NewAdobeSign())

	for _, provider := range providers {
		provider.GenerateAns()
		names = append(names, provider.GetName())
		ans = append(ans, provider.GetAns())
		sites = append(sites, strings.Join(provider.GetSearchSites(), ", "))

	}

	return ans, names, sites
}
