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

type dev struct{
	name string
	id string
	engineers []string
}

type ops struct{
	name string
	id string
	engineers []string
}

type devops struct{
	id string
	dev []string
	ops []string
}
//Global maps to access our resources by id key
engineers := make(map[string] engineer)
devs := make(map[string] devs)
ops := make(map[string] ops)
devops := make(map[string] devops)

func newDev (name string, email string) dev{
	devGroup := dev {name: name}
	devGroup.id = "TODO"
	devGroup.email = email
	devGroup.engineers = make([]engineer,1)
	return devGroup
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
