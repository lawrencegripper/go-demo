package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	requestUrl := os.Args[1]

	res, err := http.Get(requestUrl)

	if err != nil {
		fmt.Println(err)
		return
	}

	output, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(output))
}
