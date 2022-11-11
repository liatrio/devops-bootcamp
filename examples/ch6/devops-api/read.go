package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct to contain engineer as a list inside a dev resource instead of map for GET response
type dev_engineer struct {
	Name      string     `json:"name"`
	Id        string     `json:"id"`
	Engineers []engineer `json:"engineers"`
}

// struct to contain engineer as a list inside a ops resource instead of map for GET response
type ops_engineer struct {
	Name      string     `json:"name"`
	Id        string     `json:"id"`
	Engineers []engineer `json:"engineers"`
}

type devops_list struct {
	Id  string `json:"id"`
	Dev []dev  `json:"dev"`
	Ops []ops  `json:"ops"`
}

/********** Conversion function to switch maps to slices for GET response ********/
func mapToSlice(engineer_map map[string]engineer) []engineer {
	engineer_slice := make([]engineer, len(engineer_map))

	i := 0
	for _, engineer := range engineer_map {
		engineer_slice[i] = engineer
		i++
	}
	return engineer_slice
}

func devListConversion(dev_map map[string]dev) []dev_engineer {
	dev_slice := make([]dev_engineer, len(dev_map))

	i := 0
	for _, developer := range dev_map {
		dev_slice[i].Name = developer.Name
		dev_slice[i].Id = developer.Id
		dev_slice[i].Engineers = mapToSlice(developer.Engineers)
		i++
	}
	return dev_slice
}

func opsListConversion(ops_map map[string]ops) []ops_engineer {
	ops_slice := make([]ops_engineer, len(ops_map))

	i := 0
	for _, operation := range ops_map {
		ops_slice[i].Name = operation.Name
		ops_slice[i].Id = operation.Id
		ops_slice[i].Engineers = mapToSlice(operation.Engineers)
		i++
	}
	return ops_slice
}

/******************************************************/

// server GET handlers
func getEngineer(c *gin.Context) {
	engineer_slice := mapToSlice(engineers)
	c.IndentedJSON(http.StatusOK, engineer_slice)
}

func getDev(c *gin.Context) {
	dev_slice := devListConversion(developers)
	c.IndentedJSON(http.StatusOK, dev_slice)
}

func getOp(c *gin.Context) {
	ops_slice := opsListConversion(operations)
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
