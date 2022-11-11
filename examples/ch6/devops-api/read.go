package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func mapToSlice(engineer_map map[string]engineer) []engineer {
	engineer_slice := make([]engineer, len(engineer_map))

	i := 0
	for _, engineer := range engineer_map {
		engineer_slice[i] = engineer
		i++
	}
	return engineer_slice
}

// server GET handlers
func getEngineer(c *gin.Context) {
	engineer_slice := mapToSlice(engineers)
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
	devops_slice := make([]devops, len(developer_operations))
	i := 0
	for _, developer_operation := range developer_operations {
		devops_slice[i] = developer_operation
		i++
	}
	c.IndentedJSON(http.StatusOK, devops_slice)
}
