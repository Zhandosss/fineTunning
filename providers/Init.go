package providers

import "fineTunning/training"

func Init() ([][]string, []string) {
	ans := make([][]string, 0, 10)
	names := make([]string, 0, 10)

	mural := training.NewMural()
	mural.GenerateAns()
	names = append(names, mural.GetName())
	ans = append(ans, mural.GetAns())

	wrike := training.NewWrike()
	wrike.GenerateAns()
	names = append(names, wrike.GetName())
	ans = append(ans, wrike.GetAns())

	klaviyo := training.NewKlaviyo()
	klaviyo.GenerateAns()
	names = append(names, klaviyo.GetName())
	ans = append(ans, klaviyo.GetAns())

	ringCentral := training.NewRingCentral()
	ringCentral.GenerateAns()
	names = append(names, ringCentral.GetName())
	ans = append(ans, ringCentral.GetAns())

	adobeSign := training.NewAdobeSign()
	adobeSign.GenerateAns()
	names = append(names, adobeSign.GetName())
	ans = append(ans, adobeSign.GetAns())

	return ans, names
}
