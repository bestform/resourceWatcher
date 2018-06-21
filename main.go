package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	confFile, err := os.Open("urls.config")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()

	conf, err := GetResourceListFromConfig(confFile)
	if err != nil {
		panic(err)
	}

	dbFile, err := os.OpenFile(".db", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer dbFile.Close()

	err = initDb(dbFile)
	if err != nil {
		panic(err)
	}
	defer persistDb()

	for _, url := range conf {
		err = checkUrl(url)
		if err != nil {
			fmt.Println(err)
		}
	}

}

func checkUrl(url string) error {
	fmt.Println("Checking " + url)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if !HasContentForResource(url) {
		SetContentForResource(url, string(content))
		fmt.Println("New Entry for URL: " + url)
		return nil
	}
	if HasContentChanged(url, string(content)) {
		fmt.Println("Content has changed for URL: " + url)
		// @todo: if flagged 'update all' write new!
	}

	return nil
}
