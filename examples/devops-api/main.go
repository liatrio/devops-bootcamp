package main

import (
	"crypto/rand"
	"encoding/base32"
	//"errors"
	"fmt"
	"github.com/gin-gonic/gin"
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
	Name  string `json:"name"`
	Id    string `json:"id"`
	Email string `json:"email"`
}

var engineers = []engineer{

	{Name: "Bobs Burgers", Id: "12", Email: "bob@gmail.com"},
	{Name: "Alice Joice", Id: "1", Email: "alice@gmail.com"},
	{Name: "Benjamin Serrato", Id: "2", Email: "benny@gmail.com"},
}

func getEngineer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, engineers)
}

// type dev struct {
// 	name      string
// 	id        string
// 	engineers map[string]engineer
// }

// type ops struct {
// 	name      string
// 	id        string
// 	engineers map[string]engineer
// }

// type devops struct {
// 	id  string
// 	dev map[string]dev
// 	ops map[string]ops
// }

// Global maps to access our resources by id key
//var engineers = make(map[string]engineer)
// var developers = make(map[string]dev)
// var operations = make(map[string]ops)
// var developer_operations = make(map[string]devops)

// func newDevOps() (devops, error) {
// 	devOpsGroup := devops{id: "TODO"}
// 	devOpsGroup.ops = make(map[string]ops)
// 	devOpsGroup.dev = make(map[string]dev)
// 	developer_operations[devOpsGroup.id] = devOpsGroup
// 	return devOpsGroup, nil
// }

// func newDev(name string) (dev, error) {
// 	for _, value := range developers {
// 		if name == value.name {
// 			return dev{}, errors.New(" Dev group already exists ")
// 		}
// 	}
// 	devGroup := dev{name: name, id: "TODO"}
// 	devGroup.engineers = make(map[string]engineer)
// 	developers[devGroup.id] = devGroup
// 	return devGroup, nil
// }

// func newOp(name string) (ops, error) {
// 	for _, value := range operations {
// 		if name == value.name {
// 			return ops{}, errors.New(" Operations group already exists ")
// 		}
// 	}
// 	opGroup := ops{name: name, id: "TODO"}
// 	opGroup.engineers = make(map[string]engineer)
// 	operations[opGroup.id] = opGroup
// 	return opGroup, nil
// }

// func newPerson(name string, email string) (engineer, error) {
// 	for _, value := range engineers {
// 		if name == value.name {
// 			return engineer{}, errors.New(" Engineer already exists ")
// 		}
// 	}
// 	p := engineer{name: name, id: getRandId(5)}
// 	p.email = email
// 	//engineers[p.id] = p
// 	return p, nil
// }

// func addEngineerTo_Op(op_id string, engineer_id string) (bool, error) {

// 	engineer_val, exists := engineers[engineer_id]
// 	if !exists {
// 		return false, errors.New(" Engineer doesn't exist ")
// 	}
// 	dev_val, exists := developers[op_id]
// 	if !exists {
// 		return false, errors.New(" Operations group doesn't exist ")
// 	}
// 	_, exists = dev_val.engineers[engineer_id]
// 	if exists {
// 		return false, errors.New(" Engineer already exists inside specified Operations group ")
// 	}
// 	dev_val.engineers[engineer_id] = engineer_val

// 	return true, nil

// }

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

	router := gin.Default()
	router.GET("/engineers", getEngineer)

	router.Run("localhost:8080")

	fmt.Println(response)
	// functions for http handlers
}

func getRoot(write http.ResponseWriter, read *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(write, "This is the root response! Do we want each resource to be reported in a seperate response?\n")
}
