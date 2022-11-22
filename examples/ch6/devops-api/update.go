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

func updateDev(id string, newDev dev) (bool, error) {
	if newDev.Name == "" {
		return false, errors.New(" Name cannot be empty ")
	}
	//For global dev map
	dev, err := findDev_by_Id(id)
	if err != nil {
		return false, errors.New(" Doesn't exist in the developers group")
	}
	dev.Name = newDev.Name
	dev.Engineers = []*engineer{}
	for _, eng := range newDev.Engineers {
		newEngineer, err := findEngineer_by_Id(eng.Id)
		if err != nil {
			return false, errors.New(" updatedev:engineer, couldnt find engineer in global array")
		}
		dev.Engineers = append(dev.Engineers, newEngineer)
	}
	return true, nil
}

func updateOps(id string, newOp ops) (bool, error) {
	if newOp.Name == "" {
		return false, errors.New(" Name cannot be empty ")
	}
	//For global dev map
	op, err := findOp_by_Id(id)
	if err != nil {
		return false, errors.New(" Doesn't exist in the developers group")
	}
	op.Name = newOp.Name
	op.Engineers = []*engineer{}
	for _, eng := range newOp.Engineers {
		newEngineer, err := findEngineer_by_Id(eng.Id)
		if err != nil {
			return false, errors.New(" updatedev:engineer, couldnt find engineer in global array")
		}
		op.Engineers = append(op.Engineers, newEngineer)
	}
	return true, nil
}
func updateDevOps(id string, newDevOps devops) (bool, error) {
	//For global dev map
	devops, err := findDevOps_by_Id(id)
	if err != nil {
		return false, errors.New(" Doesn't exist in the developer_operations group")
	}
	devops.Dev = []*dev{}
	devops.Ops = []*ops{}
	for _, dev := range newDevOps.Dev {
		newDev, err := findDev_by_Id(dev.Id)
		if err != nil {
			return false, errors.New(" updatedevops:dev, couldnt find dev in global array")
		}
		devops.Dev = append(devops.Dev, newDev)
	}
	for _, ops := range newDevOps.Ops {
		newOp, err := findOp_by_Id(ops.Id)
		if err != nil {
			return false, errors.New(" updatedevops:ops, couldnt find op in global array")
		}
		devops.Ops = append(devops.Ops, newOp)
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

	_, err = updateDev(id, jsonData)
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

	_, err = updateOps(id, jsonData)
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

func putDevOps(c *gin.Context) {
	id := c.Param("id")
	var jsonData devops
	err := c.ShouldBindJSON(&jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = updateDevOps(id, jsonData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	devops, err := findDevOps_by_Id(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, devops)
}
