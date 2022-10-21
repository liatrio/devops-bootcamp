package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"net/http"
	"os"
)

type engineer struct {
	name  string
	id    string
	email string
}

type dev struct {
	name      string
	id        string
	engineers []string
}

type ops struct {
	name      string
	id        string
	engineers []string
}

type devops struct {
	id  string
	dev []string
	ops []string
}

// Global maps to access our resources by id key
var engineers = make(map[string]engineer)
var developers = make(map[string]dev)
var operations = make(map[string]ops)
var developer_operations = make(map[string]devops)

func newDev(name string, email string) dev {
	devGroup := dev{name: name}
	devGroup.id = "TODO"
	devGroup.engineers = make([]string, 1)
	return devGroup
}

func newPerson(name string) *engineer {
	p := engineer{name: name}
	p.id = "123"
	p.email = "asdasd@gmail.com"
	return &p
}

func getRandId(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

func verifyEmail(email string) bool {
	result, _ := regexp.MatchString("^([a-z]|[0-9])+@[a-z]+\\.[a-z]+(\\.[a-z]+)*$", email)
	return result
}

func main() {
	protocol := "http://"
	url := protocol + "www.google.com"
	response, err := http.Get(url)
	if err != nil { //if theres an error in GET request
		fmt.Printf("error in GET request: %s\nErrors: %s\n", url, err)
		os.Exit(1)
	}
	fmt.Println(response)

	randId := getRandId(15)
	//	fmt.Println("Random ID: ", randId)
	fmt.Println(engineer{"bob", randId, "hello@gmail.com"})
}
