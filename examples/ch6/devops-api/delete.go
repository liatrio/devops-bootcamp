package main

import (
	"errors"

	//"fmt"
	"github.com/gin-gonic/gin"
	//"io"
	"net/http"
)

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
	delete(operations[ops_id].Engineers, engineer_id)
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
	delete(developers[dev_id].Engineers, engineer_id)
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
	delete(developer_operations[devops_id].Dev, dev_id)

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
	delete(developer_operations[devops_id].Ops, ops_id)
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
	for dev_key, dev_val := range developers {
		for key := range dev_val.Engineers {
			if key == engineer_id {
				//delete(developer_operations[devops_key].Dev[dev_key].Engineers, engineer_id)
				delete(developers[dev_val.Id].Engineers, engineer_id)
				for devops_key, devops_val := range developer_operations {
					_, exists := devops_val.Dev[dev_key]
					if exists {
						delete(developer_operations[devops_key].Dev[dev_key].Engineers, engineer_id)
					}
				}
			}
		}
	}
	for ops_key, ops_val := range operations {
		for key := range ops_val.Engineers {
			if key == engineer_id {
				//delete(developer_operations[devops_key].Dev[dev_key].Engineers, engineer_id)
				delete(operations[ops_val.Id].Engineers, engineer_id)
				for devops_key, devops_val := range developer_operations {
					_, exists := devops_val.Ops[ops_key]
					if exists {
						delete(developer_operations[devops_key].Ops[ops_key].Engineers, engineer_id)
					}
				}
			}
		}
	}

	delete(engineers, engineer_id)
	return true, nil
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

	c.JSON(http.StatusOK, gin.H{"success": "developer operations resource deleted"})
}
