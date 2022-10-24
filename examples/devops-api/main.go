package main

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	"fmt"
	"io"
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

func newDevOps() (devops, error) {
	devOpsGroup := devops{id: "TODO"}
	devOpsGroup.ops = make(map[string]ops)
	devOpsGroup.dev = make(map[string]dev)
	developer_operations[devOpsGroup.id] = devOpsGroup
	return devOpsGroup, nil
}

func newDev(name string) (dev, error) {
	for _, value := range developers {
		if name == value.name {
			return dev{}, errors.New(" Dev group already exists ")
		}
	}
	devGroup := dev{name: name, id: "TODO"}
	devGroup.engineers = make(map[string]engineer)
	developers[devGroup.id] = devGroup
	return devGroup, nil
}

func newOp(name string) (ops, error) {
	for _, value := range operations {
		if name == value.name {
			return ops{}, errors.New(" Operations group already exists ")
		}
	}
	opGroup := ops{name: name, id: "TODO"}
	opGroup.engineers = make(map[string]engineer)
	operations[opGroup.id] = opGroup
	return opGroup, nil
}

func newPerson(name string, email string) (engineer, error) {
	for _, value := range engineers {
		if name == value.name {
			return engineer{}, errors.New(" Engineer already exists ")
		}
	}
	p := engineer{name: name, id: getRandId(5)}
	p.email = email
	engineers[p.id] = p
	return p, nil
}

func addEngineerTo_Op(op_id string, engineer_id string) (bool, error) {

	engineer_val, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	dev_val, exists := developers[op_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	_, exists = dev_val.engineers[engineer_id]
	if exists {
		return false, errors.New(" Engineer already exists inside specified Operations group ")
	}
	dev_val.engineers[engineer_id] = engineer_val

	return true, nil

}

func verifyEmail(email string) bool {
	result, _ := regexp.MatchString("^([a-z]|[0-9])+@[a-z]+\\.[a-z]+(\\.[a-z]+)*$", email)
	return result
}

//functions for http handlers

func getRoot(write http.ResponseWriter, read *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(write, "This is the root response! Do we want each resource to be reported in a seperate response?\n")
}

func main() {
	randId := getRandId(15)
	//	fmt.Println("Random ID: ", randId)
	fmt.Println(engineer{"bob", randId, "hello@gmail.com"})

	http.HandleFunc("/", getRoot)
	error := http.ListenAndServe(":8000", nil)
	if errors.Is(error, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if error != nil {
		fmt.Printf("error starting server: %s\n", error)
		os.Exit(1)
	}
}
