package main

import (
	"io"
	"io/ioutil"
	"strings"
)

// GetResourceListFromConfig will read a list of URLs separated by new lines
func GetResourceListFromConfig(r io.Reader) ([]string, error) {
	urls := make([]string, 0)
	conf, err := ioutil.ReadAll(r)
	if err != nil {
		return urls, err
	}

	rawUrls := strings.Split(string(conf), "\n")
	var trimmedUrl string
	for _, url := range rawUrls {
		trimmedUrl = strings.Trim(url, " ")
		if trimmedUrl != "" {
			urls = append(urls, trimmedUrl)
		}
	}

	return urls, nil
}
