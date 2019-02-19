package main

import ( 
		"fmt"
		"net/http"
		"io/ioutil"
		)
func main(){
	url := "https://en.wikipedia.org/wiki/Main_Page"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	fmt.Printf("%s\n", html)

}

