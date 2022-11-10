package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"net/http"
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
	engineer_val, exists := engineers[engineer_id]
	if exists {
		engineer_val.Email = email
		engineer_val.Name = name
		engineers[engineer_id] = engineer_val
	} else {
		return false, errors.New(" Engineer doesn't exist ")
	}

	for dev_key, dev_val := range developers {
		for engineer_key, engineer_val := range dev_val.Engineers {
			if engineer_key == engineer_id {
				engineer_val.Email = email
				engineer_val.Name = name
				developers[dev_key].Engineers[engineer_key] = engineer_val
				for devops_key, devops_val := range developer_operations {
					_, exists := devops_val.Dev[dev_key]
					if exists {
						developer_operations[devops_key].Dev[dev_key].Engineers[engineer_key] = engineer_val
					}
				}
			}
		}
	}
	for ops_key, ops_val := range operations {
		for engineer_key, engineer_val := range ops_val.Engineers {
			if engineer_key == engineer_id {
				engineer_val.Email = email
				engineer_val.Name = name
				operations[ops_key].Engineers[engineer_key] = engineer_val
				for devops_key, devops_val := range developer_operations {
					_, exists := devops_val.Ops[ops_key]
					if exists {
						developer_operations[devops_key].Ops[ops_key].Engineers[engineer_key] = engineer_val
					}
				}
			}
		}
	}
	return true, nil
}

func updateDev(dev_id string, name string) (bool, error) {
	if name == "" {
		return false, errors.New(" Name cannot be empty ")
	}
	//For global dev map
	dev_map_val, exists := developers[dev_id]
	if exists {
		dev_map_val.Name = name
		developers[dev_id] = dev_map_val
	} else {
		return false, errors.New(" Doesn't exist in the developers group")
	}
	//For updating the values for developers inside global developer_operations map
	for devops_key, devops_val := range developer_operations {
		for dev_key, dev_val := range devops_val.Dev {
			if dev_key == dev_id {
				dev_val.Name = name
				developer_operations[devops_key].Dev[dev_key] = dev_val
				//developers[dev_key] = dev_val
			}
		}
	}
	return true, nil
}

func updateOps(ops_id string, name string) (bool, error) {
	if name == "" {
		return false, errors.New(" Name cannot be empty ")
	}
	//For global dev map
	op_map_val, exists := operations[ops_id]
	if exists {
		op_map_val.Name = name
		operations[ops_id] = op_map_val
	} else {
		return false, errors.New(" Doesn't exist in the operations group")
	}
	//For updating the values for developers inside global developer_operations map
	for devops_key, devops_val := range developer_operations {
		for ops_key, ops_val := range devops_val.Ops {
			if ops_key == ops_id {
				ops_val.Name = name
				developer_operations[devops_key].Ops[ops_key] = ops_val
				//operations[ops_key] = ops_val
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
	c.IndentedJSON(http.StatusOK, engineers[id])
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
	c.IndentedJSON(http.StatusOK, developers[id])
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
	c.IndentedJSON(http.StatusOK, operations[id])
}
