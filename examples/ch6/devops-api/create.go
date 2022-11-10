package main

import (
	"errors"

	//"fmt"
	"github.com/gin-gonic/gin"
	//"io"
	"net/http"
)

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
	if !verifyEmail(email) {
		return engineer{}, errors.New(" Email is invalid ")
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
		return false, errors.New(" Developer Operations group doesn't exist ")
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
		return false, errors.New(" Operations group doesn't exist ")
	}
	devops_val, exists := developer_operations[devops_id]
	if !exists {
		return false, errors.New(" Developer Operations group doesn't exist ")
	}
	_, exists = devops_val.Ops[op_id]
	if exists {
		return false, errors.New(" Operations group already exists inside specified Developer Operations group ")
	}
	developer_operations[devops_id].Ops[op_id] = op_val

	return true, nil

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

func postDevEngineer(c *gin.Context) {
	id := c.Param("id") //dev id
	var jsonData engineer
	err := c.ShouldBindJSON(&jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = addEngineerTo_Dev(id, jsonData.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, developers[id])
}

func postOpEngineer(c *gin.Context) {
	id := c.Param("id") //op id
	var jsonData engineer
	err := c.ShouldBindJSON(&jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = addEngineerTo_Op(id, jsonData.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, operations[id])
}

func postDevOpsDev(c *gin.Context) {
	id := c.Param("id") //devops id
	var jsonData dev

	err := c.ShouldBindJSON(&jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = addDevTo_DevOps(id, jsonData.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, developer_operations[id])

}

func postDevOpsOp(c *gin.Context) {
	id := c.Param("id") //devops id
	var jsonData ops

	err := c.ShouldBindJSON(&jsonData)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = addOpTo_DevOps(id, jsonData.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, developer_operations[id])

}
