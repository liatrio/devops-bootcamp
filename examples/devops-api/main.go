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
	Name  string `json:"name" binding:"required"`
	Id    string `json:"id"`
	Email string `json:"email" binding:"required"`
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

func newDevOps() (devops, error) {
	devOpsGroup := devops{Id: getRandId(5)}
	devOpsGroup.Ops = make(map[string]ops)
	devOpsGroup.Dev = make(map[string]dev)
	developer_operations[devOpsGroup.Id] = devOpsGroup
	return devOpsGroup, nil
}

func newDev(name string) (dev, error) {
	if name == "" {
		return dev{}, errors.New(" Name cannot be empty ")
	}
	for _, value := range developers {
		if name == value.Name {
			return dev{}, errors.New(" Dev group already exists ")
		}
	}
	devGroup := dev{Name: name, Id: getRandId(5)}
	devGroup.Engineers = make(map[string]engineer)
	developers[devGroup.Id] = devGroup
	return devGroup, nil
}

func newOp(name string) (ops, error) {
	if name == "" {
		return ops{}, errors.New(" Name cannot be empty ")
	}
	for _, value := range operations {
		if name == value.Name {
			return ops{}, errors.New(" Operations group already exists ")
		}
	}
	opGroup := ops{Name: name, Id: getRandId(5)}
	opGroup.Engineers = make(map[string]engineer)
	operations[opGroup.Id] = opGroup
	return opGroup, nil
}

func newEngineer(name string, email string) (engineer, error) {
	if name == "" {
		return engineer{}, errors.New(" Name cannot be empty ")
	}
	for _, value := range engineers {
		if name == value.Name {
			return engineer{}, errors.New(" Engineer already exists ")
		}
	}
	p := engineer{Name: name, Id: getRandId(5)}
	p.Email = email
	engineers[p.Id] = p
	return p, nil
}

// functions to add resources to other resources//
func addEngineerTo_Op(ops_id string, engineer_id string) (bool, error) {

	engineer_val, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	op_val, exists := operations[ops_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	_, exists = op_val.Engineers[engineer_id]
	if exists {
		return false, errors.New(" Engineer already exists inside specified Operations group ")
	}
	operations[ops_id].Engineers[engineer_id] = engineer_val
	for devops_key, devops_val := range developer_operations {
		for ops_key := range devops_val.Ops {
			if ops_key == ops_id {
				developer_operations[devops_key].Ops[ops_id].Engineers[engineer_id] = engineer_val
			}
		}
	}

	return true, nil

}

func addEngineerTo_Dev(dev_id string, engineer_id string) (bool, error) {

	engineer_val, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	dev_val, exists := developers[dev_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	_, exists = dev_val.Engineers[engineer_id]
	if exists {
		return false, errors.New(" Engineer already exists inside specified Operations group ")
	}
	developers[dev_id].Engineers[engineer_id] = engineer_val
	for devops_key, devops_val := range developer_operations {
		for dev_key := range devops_val.Dev {
			if dev_key == dev_id {
				developer_operations[devops_key].Dev[dev_id].Engineers[engineer_id] = engineer_val
			}
		}
	}

	return true, nil

}

func addDevTo_DevOps(devops_id string, dev_id string) (bool, error) {

	dev_val, exists := developers[dev_id]
	if !exists {
		return false, errors.New(" Developer group doesn't exist ")
	}
	devops_val, exists := developer_operations[devops_id]
	if !exists {
		return false, errors.New(" Developer Oerations group doesn't exist ")
	}
	_, exists = devops_val.Dev[dev_id]
	if exists {
		return false, errors.New(" Developer already exists inside specified Developer Operations group ")
	}
	developer_operations[devops_id].Dev[dev_id] = dev_val

	return true, nil

}

func addOpTo_DevOps(devops_id string, op_id string) (bool, error) {

	op_val, exists := operations[op_id]
	if !exists {
		return false, errors.New(" Developer group doesn't exist ")
	}
	devops_val, exists := developer_operations[devops_id]
	if !exists {
		return false, errors.New(" Developer Oerations group doesn't exist ")
	}
	_, exists = devops_val.Ops[op_id]
	if exists {
		return false, errors.New(" Developer already exists inside specified Developer Operations group ")
	}
	developer_operations[devops_id].Ops[op_id] = op_val

	return true, nil

}

// *********************************************//
// functions to delete resources from other resources//
func deleteEngineerFrom_Op(ops_id string, engineer_id string) (bool, error) {

	_, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	op_val, exists := operations[ops_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	_, exists = op_val.Engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exists inside specified Operations group ")
	}
	delete(op_val.Engineers, engineer_id)
	for devops_key, devops_val := range developer_operations {
		for key := range devops_val.Ops {
			if key == ops_id {
				delete(developer_operations[devops_key].Ops[ops_id].Engineers, engineer_id)
			}
		}
	}

	return true, nil

}

func deleteEngineerFrom_Dev(dev_id string, engineer_id string) (bool, error) {

	_, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	dev_val, exists := developers[dev_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	_, exists = dev_val.Engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesnt exists inside specified Operations group ")
	}
	delete(dev_val.Engineers, engineer_id)
	for devops_key, devops_val := range developer_operations {
		for key := range devops_val.Dev {
			if key == dev_id {
				delete(developer_operations[devops_key].Dev[dev_id].Engineers, engineer_id)
			}
		}
	}

	return true, nil

}

func deleteDevFrom_DevOps(devops_id string, dev_id string) (bool, error) {

	_, exists := developers[dev_id]
	if !exists {
		return false, errors.New(" Developer group doesn't exist ")
	}
	devops_val, exists := developer_operations[devops_id]
	if !exists {
		return false, errors.New(" Developer Oerations group doesn't exist ")
	}
	_, exists = devops_val.Dev[dev_id]
	if !exists {
		return false, errors.New(" Developer doesn't exists inside specified Developer Operations group ")
	}
	delete(devops_val.Dev, dev_id)

	return true, nil

}

func deleteOpFrom_DevOps(devops_id string, ops_id string) (bool, error) {

	_, exists := developers[ops_id]
	if !exists {
		return false, errors.New(" Developer group doesn't exist ")
	}
	devops_val, exists := developer_operations[devops_id]
	if !exists {
		return false, errors.New(" Developer Oerations group doesn't exist ")
	}
	_, exists = devops_val.Dev[ops_id]
	if !exists {
		return false, errors.New(" Developer doesn't exists inside specified Developer Operations group ")
	}
	delete(devops_val.Ops, ops_id)
	return true, nil

}

// **************************************************//
// functions to delete resources//
func deleteDevOps(devops_id string) (bool, error) {
	_, exists := developer_operations[devops_id]
	if !exists {
		return false, errors.New(" Developer Oerations group doesn't exist ")
	}
	delete(developer_operations, devops_id)
	return true, nil
}

func deleteDev(dev_id string) (bool, error) {
	_, exists := developers[dev_id]
	if !exists {
		return false, errors.New(" Developer group doesn't exist ")
	}
	for devops_key, devops_val := range developer_operations {
		for key, _ := range devops_val.Dev {
			if key == dev_id {
				delete(developer_operations[devops_key].Dev, dev_id)
			}
		}
	}
	delete(developers, dev_id)
	return true, nil
}

func deleteOp(ops_id string) (bool, error) {
	_, exists := operations[ops_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	for devops_key, devops_val := range developer_operations {
		for key, _ := range devops_val.Ops {
			if key == ops_id {
				delete(developer_operations[devops_key].Ops, ops_id)
			}
		}
	}
	delete(operations, ops_id)
	return true, nil
}

func deleteEngineer(engineer_id string) (bool, error) {
	_, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	for devops_key, devops_val := range developer_operations {
		for dev_key, dev_val := range devops_val.Dev {
			for key, _ := range dev_val.Engineers {
				if key == engineer_id {
					delete(developer_operations[devops_key].Dev[dev_key].Engineers, engineer_id)
					delete(developers[dev_val.Id].Engineers, engineer_id)
				}
			}
		}
	}
	for devops_key, devops_val := range developer_operations {
		for ops_key, ops_val := range devops_val.Ops {
			for key, _ := range ops_val.Engineers {
				if key == engineer_id {
					delete(developer_operations[devops_key].Ops[ops_key].Engineers, engineer_id)
					delete(operations[ops_val.Id].Engineers, engineer_id)
				}
			}
		}
	}
	delete(engineers, engineer_id)
	return true, nil
}

// *****************************//

// functions to update resources//
func updateEngineer(engineer_id string, name string, email string) (bool, error) {
	//For updating global engineers map
	engineer_val, exists := engineers[engineer_id]
	if exists {
		engineer_val.Email = email
		engineer_val.Name = name
		engineers[engineer_id] = engineer_val
	} else {
		return false, errors.New(" Engineer doesn't exist ")
	}
	//For updating the values for developers inside global developer_operations map
	for devops_key, devops_val := range developer_operations {
		for dev_key, dev_val := range devops_val.Dev {
			for key, engineer_val := range dev_val.Engineers {
				if key == engineer_id {
					engineer_val.Email = email
					engineer_val.Name = name
					developers[dev_key].Engineers[key] = engineer_val
					developer_operations[devops_key].Dev[dev_key].Engineers[key] = engineer_val
				}
			}
		}
	}

	//For updating the values for operations inside global developer_operations map
	for devops_key, devops_val := range developer_operations {
		for ops_key, ops_val := range devops_val.Ops {
			for key, engineer_val := range ops_val.Engineers {
				if key == engineer_id {
					engineer_val.Email = email
					engineer_val.Name = name
					operations[ops_key].Engineers[key] = engineer_val
					developer_operations[devops_key].Ops[ops_key].Engineers[key] = engineer_val
				}
			}
		}
	}
	return true, nil
}

func updateDev(dev_id string, name string) (bool, error) {
	//For global dev map
	dev_map_val, exists := developers[dev_id]
	if exists {
		dev_map_val.Name = name
	} else {
		return false, errors.New(" Doesn't exist in the developers group")
	}
	//For updating the values for developers inside global developer_operations map
	for devops_key, devops_val := range developer_operations {
		for dev_key, dev_val := range devops_val.Dev {
			if dev_key == dev_id {
				dev_val.Name = name
				developer_operations[devops_key].Dev[dev_key] = dev_val
				developers[dev_key] = dev_val
			}
		}
	}
	return true, nil
}

func updateOps(ops_id string, name string) (bool, error) {
	//For global dev map
	op_map_val, exists := operations[ops_id]
	if exists {
		op_map_val.Name = name
	} else {
		return false, errors.New(" Doesn't exist in the operations group")
	}
	//For updating the values for developers inside global developer_operations map
	for devops_key, devops_val := range developer_operations {
		for ops_key, ops_val := range devops_val.Ops {
			if ops_key == ops_id {
				ops_val.Name = name
				developer_operations[devops_key].Ops[ops_key] = ops_val
				operations[ops_key] = ops_val
			}
		}
	}
	return true, nil
}

/*
func updateDevOps(devops_id string) (bool, error) {
	//TODO
	return true, nil
}
*/
//*****************************//

func verifyEmail(email string) bool {
	result, _ := regexp.MatchString("^([a-z]|[0-9])+@[a-z]+\\.[a-z]+(\\.[a-z]+)*$", email)
	return result
}

// server GET handlers
func getEngineer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, engineers)
}

func getDev(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, developers)
}

func getOp(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, operations)
}

func getDevOps(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, developer_operations)
}

// server POST handlers
func postEngineer(c *gin.Context) {
	var jsonData engineer    //object that gets name and email from POST request
	var curEngineer engineer //object recieved from newEngineer

	err := c.BindJSON(&jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curEngineer, err = newEngineer(jsonData.Name, jsonData.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, engineers[curEngineer.Id])
}

func postDev(c *gin.Context) {
	var jsonData dev //object that gets dev data from POST request
	var curDev dev   //object recieved from newDev

	err := c.BindJSON(&jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curDev, err = newDev(jsonData.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, developers[curDev.Id])
}

func postOp(c *gin.Context) {
	var jsonData ops //object that gets dev data from POST request
	var curOp ops    //object recieved from newOp

	err := c.BindJSON(&jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	curOp, err = newOp(jsonData.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, operations[curOp.Id])

}

func postDevOps(c *gin.Context) {
	curDevOps, err := newDevOps()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, developer_operations[curDevOps.Id])
}

// server PUT handler
func putEngineer(c *gin.Context) {
	id := c.Param("id")
	var jsonData engineer
	//	var result bool
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = updateEngineer(id, jsonData.Name, jsonData.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, engineers[id])
}

// server DELETE handler
func deleteRequestEngineer(c *gin.Context) {
	id := c.Param("id")
	_, err := deleteEngineer(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "engineer resource deleted"})
}

func deleteRequestDev(c *gin.Context) {
	id := c.Param("id")
	_, err := deleteDev(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "developer resource deleted"})
}

func deleteRequestOp(c *gin.Context) {
	id := c.Param("id")
	_, err := deleteOp(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "operations resource deleted"})
}

func deleteRequestDevOps(c *gin.Context) {
	id := c.Param("id")
	_, err := deleteDevOps(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "operations resource deleted"})
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
	//PUT routes
	router.PUT("/engineers/:id", putEngineer)
	//DELETE routes
	router.DELETE("/engineers/:id", deleteRequestEngineer)
	router.DELETE("/dev/:id", deleteRequestDev)
	router.DELETE("/op/:id", deleteRequestOp)
	router.DELETE("/devops/:id", deleteRequestDevOps)
	//runs server
	router.Run(":8080")
}
