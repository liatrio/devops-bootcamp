package main

import (
	"crypto/rand"
	"encoding/base32"
	"errors"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"io"
	"net/http"
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

func newEngineer(name string, email string) (engineer, error) {
	for _, value := range engineers {
		if name == value.Name {
			return engineer{}, errors.New(" Engineer already exists ")
		}
	}
	p := engineer{Name: name, Id: getRandId(5)}
	p.Email = email
	//engineers[p.id] = p
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

// server GET handlers
func getEngineer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, exengineers)
}

// server POST handlers
func postEngineer(c *gin.Context) {
	var jsonData engineer    //object that gets name and email from POST request
	var curEngineer engineer //object recieved from newEngineer

	err := c.BindJSON(&jsonData)

	if err != nil {
		return
	}

	curEngineer, err = newEngineer(jsonData.Name, jsonData.Email)
	engineers[curEngineer.Id] = curEngineer
	c.IndentedJSON(http.StatusOK, engineers)
}

func main() {
	router := gin.Default()

	//GET routes
	router.GET("/engineers", getEngineer)
	//POST routes
	router.POST("/engineers", postEngineer)

	//runs server
	router.Run("localhost:8080")
}
