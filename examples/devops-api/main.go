package main

import(
	"os"
	"fmt"
	"net/http"
	"crypto/rand"
	"encoding/base32"
)


type engineer struct{
	name string
	id string
	email string
}

func newPerson (name string) *engineer {
	p := engineer {name: name}
	p.id := "123"
	p.email = "asdasd@gmail.com"
	return &p
}

func main() {
	protocol := "http://"
	url := protocol + "www.google.com"
	response, err := http.Get(url)
	if err != nil { //if theres an error in GET request
		fmt.Printf("error in GET request: %s\nErrors: %s\n",url, err)
		os.Exit(1)
	}
	fmt.Println(response)

	randId := getRandId(15)
//	fmt.Println("Random ID: ", randId)
	fmt.Println(engineer{"bob",randId ,"hello@gmail.com"})
}

func getRandId(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}
