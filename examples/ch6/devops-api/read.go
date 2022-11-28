package main

import (
	"errors"
	"net/http"

	"github.com/baserrato/devops-resource"
	"github.com/gin-gonic/gin"
)

/********** Conversion function to switch maps to slices for GET response ********/
/*
func mapToSlice(engineer_map map[string]engineer) []engineer {
	engineer_slice := make([]engineer, len(engineer_map))

	i := 0
	for _, engineer := range engineer_map {
		engineer_slice[i] = engineer
		i++
	}
	return engineer_slice
}

func devToDevEngineer(developer dev) dev_engineer {
	var developer_engineer dev_engineer
	developer_engineer.Name = developer.Name
	developer_engineer.Id = developer.Id
	developer_engineer.Engineers = mapToSlice(developer.Engineers)

	return developer_engineer
}

func opsToOpsEngineer(operation ops) ops_engineer {
	var operation_engineer ops_engineer
	operation_engineer.Name = operation.Name
	operation_engineer.Id = operation.Id
	operation_engineer.Engineers = mapToSlice(operation.Engineers)

	return operation_engineer
}

func devListConversion(dev_map map[string]dev) []dev_engineer {
	dev_slice := make([]dev_engineer, len(dev_map))

	i := 0
	for _, developer := range dev_map {
		dev_slice[i] = devToDevEngineer(developer)
		i++
	}
	return dev_slice
}

func opsListConversion(ops_map map[string]ops) []ops_engineer {
	ops_slice := make([]ops_engineer, len(ops_map))

	i := 0
	for _, operation := range ops_map {
		ops_slice[i] = opsToOpsEngineer(operation)
		i++
	}
	return ops_slice
}

func devopsListConversion(devops_map map[string]devops) []devops_list {
	devops_slice := make([]devops_list, len(devops_map))
	i := 0
	for _, developer_operation := range devops_map {
		devops_slice[i].Id = developer_operation.Id
		devops_slice[i].Dev = devListConversion(developer_operation.Dev)
		devops_slice[i].Ops = opsListConversion(developer_operation.Ops)
		i++
	}
	return devops_slice
}
*/
/******************************************************/

func findEngineer_by_Name(engineer_name string) (*devops_resource.Engineer, error) {
	for _, newEngineer := range engineers {
		if newEngineer.Name == engineer_name {
			return newEngineer, nil
		}
	}
	return nil, errors.New(" no engineer with the name ")
}

func findEngineer_by_Email(engineer_email string) (*devops_resource.Engineer, error) {
	for _, newEngineer := range engineers {
		if newEngineer.Email == engineer_email {
			return newEngineer, nil
		}
	}
	return nil, errors.New(" no engineer with that email ")
}

func findDev_by_Name(dev_name string) (*devops_resource.Dev, error) {
	for _, newDev := range developers {
		if newDev.Name == dev_name {
			return newDev, nil
		}
	}
	return nil, errors.New(" no dev group with that name ")
}

func findOps_by_Name(ops_name string) (*devops_resource.Ops, error) {
	for _, newOps := range operations {
		if newOps.Name == ops_name {
			return newOps, nil
		}
	}
	return nil, errors.New(" no ops group with that name ")
}

func getSpecificEngineerById(c *gin.Context) {
	id := c.Param("id")

	engineer, err := findEngineer_by_Id(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, engineer)
}

func getSpecificEngineerByName(c *gin.Context) {
	name := c.Param("name")

	engineer, err := findEngineer_by_Name(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, engineer)
}

func getSpecificEngineerByEmail(c *gin.Context) {
	email := c.Param("email")

	engineer, err := findEngineer_by_Email(email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, engineer)
}

func getSpecificDevById(c *gin.Context) {
	id := c.Param("id")

	dev, err := findDev_by_Id(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//dev_slice := devToDevEngineer(dev)
	c.IndentedJSON(http.StatusOK, dev)
}

func getSpecificDevByName(c *gin.Context) {
	name := c.Param("name")

	dev, err := findDev_by_Name(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, dev)
}

func getSpecificOpsById(c *gin.Context) {
	id := c.Param("id")

	ops, err := findOp_by_Id(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//ops_slice := opsToOpsEngineer(ops)
	c.IndentedJSON(http.StatusOK, ops)

}

func getSpecificOpsByName(c *gin.Context) {
	name := c.Param("name")

	ops, err := findOps_by_Name(name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusOK, ops)
}

func getSpecificDevOpsById(c *gin.Context) {
	id := c.Param("id")

	devops, err := findDevOps_by_Id(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//ops_slice := opsToOpsEngineer(ops)
	c.IndentedJSON(http.StatusOK, devops)

}

func getEngineer(c *gin.Context) {
	//engineer_slice := mapToSlice(engineers)
	c.IndentedJSON(http.StatusOK, engineers)
}

func getDev(c *gin.Context) {
	//dev_slice := devListConversion(developers)
	c.IndentedJSON(http.StatusOK, developers)
}

func getOp(c *gin.Context) {
	//ops_slice := opsListConversion(operations)
	c.IndentedJSON(http.StatusOK, operations)
}

func getDevOps(c *gin.Context) {
	//devops_slice := devopsListConversion(developer_operations)
	c.IndentedJSON(http.StatusOK, developer_operations)
}
