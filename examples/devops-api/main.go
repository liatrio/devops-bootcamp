package main

import (
	"crypto/rand"
	"encoding/base32"
	"fmt"
	"net/http"
	"os"
	"regexp"
)

func getRandId(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

type engineer struct {
	name  string
	id    string
	email string
}

type dev struct {
	name      string
	id        string
	engineers map[string]engineer
}

type ops struct {
	name      string
	id        string
	engineers map[string]engineer
}

type devops struct {
	id  string
	dev map[string]dev
	ops map[string]ops
}

// Global maps to access our resources by id key
var engineers = make(map[string]engineer)
var developers = make(map[string]dev)
var operations = make(map[string]ops)
var developer_operations = make(map[string]devops)

func newDev(name string) dev {
	for _, value := range developers {
		if name == value.name {
			//Throw warning
		}
	}
	devGroup := dev{name: name, id: "TODO"}
	devGroup.engineers = make(map[string]engineer)
	return devGroup
}

func newOp(name string) dev {
	for _, value := range operations {
		if name == value.name {
			//Throw warning
		}
	}
	opGroup := dev{name: name, id: "TODO"}
	opGroup.engineers = make(map[string]engineer)
	return opGroup
}

func newPerson(name string) engineer {
	for _, value := range engineers {
		if name == value.name {
			//Throw warning
		}
	}
	p := engineer{name: name, id: getRandId(5)}
	p.email = "asdasd@gmail.com"
	return p
}

func addEngineerTo_Op(op_id string, engineer_id string) bool {

	engineer_val, exists := engineers[engineer_id]
	if exists {
		//Throw error, no engineer
	}
	dev_val, exists := developers[op_id]
	if !exists {
		//Throw error, no operations team
	}
	_, exists = dev_val.engineers[engineer_id]
	if exists {
		//Throw error, engineer already exists in operations team
	}
	dev_val.engineers[engineer_id] = engineer_val

	return false

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
