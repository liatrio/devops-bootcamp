package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// functions to update resources//
func updateEngineer(engineer_id string, name string, email string) (bool, error) {
	if !verifyEmail(email) {
		return false, errors.New(" Email is invalid ")
	}
	if name == "" {
		return false, errors.New(" Name cannot be empty ")
	}
	//For updating global engineers map
	engineer, err := findEngineer_by_Id(engineer_id)
	if err == nil {
		engineer.Email = email
		engineer.Name = name
	} else {
		return false, errors.New(" Engineer doesn't exist ")
	}
	return true, nil
}

func updateDev(dev_id string, name string) (bool, error) {
	if name == "" {
		return false, errors.New(" Name cannot be empty ")
	}
	//For global dev map
	dev, err := findDev_by_Id(dev_id)
	if err == nil {
		dev.Name = name
	} else {
		return false, errors.New(" Doesn't exist in the developers group")
	}
	return true, nil
}

func updateOps(op_id string, name string) (bool, error) {
	if name == "" {
		return false, errors.New(" Name cannot be empty ")
	}
	//For global dev map
	op, err := findOp_by_Id(op_id)
	if err == nil {
		op.Name = name
	} else {
		return false, errors.New(" Doesn't exist in the operations group")
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

// server PUT handler
func putEngineer(c *gin.Context) {
	id := c.Param("id")
	var jsonData engineer
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
	engineer, err := findEngineer_by_Id(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, engineer)
}

func putDev(c *gin.Context) {
	id := c.Param("id")
	var jsonData dev
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = updateDev(id, jsonData.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dev, err := findDev_by_Id(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, dev)
}

func putOp(c *gin.Context) {
	id := c.Param("id")
	var jsonData ops
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = updateOps(id, jsonData.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	op, err := findOp_by_Id(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, op)
}
