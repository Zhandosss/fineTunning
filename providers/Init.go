package providers

import (
	"fineTunning/training"
	"strings"
)

func Init() ([][]string, []string, []string) {
	ans := make([][]string, 0, 10)
	names := make([]string, 0, 10)
	sites := make([]string, 0, 10)

	mural := training.NewMural()
	mural.GenerateAns()
	names = append(names, mural.GetName())
	ans = append(ans, mural.GetAns())
	sites = append(sites, strings.Join(mural.SearchSites, ", "))

	wrike := training.NewWrike()
	wrike.GenerateAns()
	names = append(names, wrike.GetName())
	ans = append(ans, wrike.GetAns())
	sites = append(sites, strings.Join(wrike.SearchSites, ", "))

	klaviyo := training.NewKlaviyo()
	klaviyo.GenerateAns()
	names = append(names, klaviyo.GetName())
	ans = append(ans, klaviyo.GetAns())
	sites = append(sites, strings.Join(klaviyo.SearchSites, ", "))

	ringCentral := training.NewRingCentral()
	ringCentral.GenerateAns()
	names = append(names, ringCentral.GetName())
	ans = append(ans, ringCentral.GetAns())
	sites = append(sites, strings.Join(ringCentral.SearchSites, ", "))

	adobeSign := training.NewAdobeSign()
	adobeSign.GenerateAns()
	names = append(names, adobeSign.GetName())
	ans = append(ans, adobeSign.GetAns())
	sites = append(sites, strings.Join(adobeSign.SearchSites, ", "))

	return ans, names, sites
}
