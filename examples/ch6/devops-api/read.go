package main

import (
	"net/http"

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

func getSpecificEngineer(c *gin.Context) {
	id := c.Param("id")

	engineer, err := findEngineer_by_Id(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Engineer resource does not exist"})
		return
	}

	c.IndentedJSON(http.StatusOK, engineer)
}

func getSpecificDev(c *gin.Context) {
	id := c.Param("id")

	dev, err := findDev_by_Id(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Developer resource does not exist"})
		return
	}

	//dev_slice := devToDevEngineer(dev)
	c.IndentedJSON(http.StatusOK, dev)
}

func getSpecificOps(c *gin.Context) {
	id := c.Param("id")

	ops, err := findOp_by_Id(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Operations resource does not exist"})
		return
	}

	//ops_slice := opsToOpsEngineer(ops)
	c.IndentedJSON(http.StatusOK, ops)

}

func getSpecificDevOps(c *gin.Context) {
	id := c.Param("id")

	devops, err := findDevOps_by_Id(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "DevOps resource does not exist"})
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
