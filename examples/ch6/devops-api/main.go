package main

import (
	"crypto/rand"
	"encoding/base32"

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

// server GET handlers
func getEngineer(c *gin.Context) {
	engineer_slice := make([]engineer, len(engineers))

	i := 0
	for _, engineer := range engineers {
		engineer_slice[i] = engineer
		i++
	}
	c.IndentedJSON(http.StatusOK, engineer_slice)
}

func getDev(c *gin.Context) {
	dev_slice := make([]dev, len(developers))

	i := 0
	for _, developer := range developers {
		dev_slice[i] = developer
		i++
	}
	c.IndentedJSON(http.StatusOK, dev_slice)
}

func getOp(c *gin.Context) {
	ops_slice := make([]ops, len(operations))
	i := 0
	for _, operation := range operations {
		ops_slice[i] = operation
		i++
	}
	c.IndentedJSON(http.StatusOK, ops_slice)
}

func getDevOps(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, developer_operations)
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
	router.POST("/dev/:id", postDevEngineer)
	router.POST("/op/:id", postOpEngineer)
	router.POST("/op", postOp)
	router.POST("/devops", postDevOps)
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
