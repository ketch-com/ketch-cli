package assets

import "io/ioutil"

func GetAsset(url string) string {
	f, err := Assets.Open(url)
	if err != nil {
		return ""
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return ""
	}

	return string(b)
}
