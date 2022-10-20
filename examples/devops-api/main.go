package main

import(
    "os"
    "fmt"
    "net/http"
)

func main() {
	protocol := "http://"
	url := protocol + "www.google.com"
	response, err := http.Get(url)
	if err != nil { //if theres an error in GET request
	    fmt.Printf("error in GET request: %s\nErrors: %s\n",url, err)
	    os.Exit(1)
	}
	fmt.Println(response)
}
