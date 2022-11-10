package main

import (
	"crypto/rand"
	"encoding/base32"
	"regexp"

	"github.com/gin-gonic/gin"
)

type engineer struct {
	Name  string `json:"name"`
	Id    string `json:"id"`
	Email string `json:"email"`
}

type dev struct {
	Name      string              `json:"name"`
	Id        string              `json:"id"`
	Engineers map[string]engineer `json:"engineers"`
}

type ops struct {
	Name      string              `json:"name"`
	Id        string              `json:"id"`
	Engineers map[string]engineer `json:"engineers"`
}

type devops struct {
	Id  string         `json:"id"`
	Dev map[string]dev `json:"dev"`
	Ops map[string]ops `json:"ops"`
}

// Global maps to access our resources by id key
var engineers = make(map[string]engineer)
var developers = make(map[string]dev)
var operations = make(map[string]ops)
var developer_operations = make(map[string]devops)

func verifyEmail(email string) bool {
	result, _ := regexp.MatchString("^([a-zA-Z]|[0-9])+@[a-z]+\\.[a-z]+(\\.[a-z]+)*$", email)
	return result
}

func getRandId(length int) string {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}
	return base32.StdEncoding.EncodeToString(randomBytes)[:length]
}

func main() {
	router := gin.Default()

	//GET routes
	router.GET("/engineers", getEngineer)
	router.GET("/dev", getDev)
	router.GET("/op", getOp)
	router.GET("/devops", getDevOps)
	//POST routes
	router.POST("/engineers", postEngineer)
	router.POST("/dev", postDev)
	router.POST("/op", postOp)
	router.POST("/devops", postDevOps)
	router.POST("/dev/:id", postDevEngineer)
	router.POST("/op/:id", postOpEngineer)
	router.POST("/devops/dev/:id", postDevOpsDev)
	router.POST("/devops/op/:id", postDevOpsOp)
	//PUT routes
	router.PUT("/engineers/:id", putEngineer)
	router.PUT("/dev/:id", putDev)
	router.PUT("/op/:id", putOp)
	//DELETE routes
	router.DELETE("/engineers/:id", deleteRequestEngineer)
	router.DELETE("/dev/:id", deleteRequestDev)
	router.DELETE("/op/:id", deleteRequestOp)
	router.DELETE("/devops/:id", deleteRequestDevOps)

	//runs server
	router.Run(":8080")
}
