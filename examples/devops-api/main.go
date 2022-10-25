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
	name      string
	id        string
	engineers map[string]engineer
}

type ops struct {
	name      string
	id        string
	engineers map[string]engineer
}

type devops struct {
	id  string
	dev map[string]dev
	ops map[string]ops
}

// Global maps to access our resources by id key
var engineers = make(map[string]engineer)
var developers = make(map[string]dev)
var operations = make(map[string]ops)
var developer_operations = make(map[string]devops)

func newDevOps() (devops, error) {
	devOpsGroup := devops{id: "TODO"}
	devOpsGroup.ops = make(map[string]ops)
	devOpsGroup.dev = make(map[string]dev)
	developer_operations[devOpsGroup.id] = devOpsGroup
	return devOpsGroup, nil
}

func newDev(name string) (dev, error) {
	for _, value := range developers {
		if name == value.name {
			return dev{}, errors.New(" Dev group already exists ")
		}
	}
	devGroup := dev{name: name, id: "TODO"}
	devGroup.engineers = make(map[string]engineer)
	developers[devGroup.id] = devGroup
	return devGroup, nil
}

func newOp(name string) (ops, error) {
	for _, value := range operations {
		if name == value.name {
			return ops{}, errors.New(" Operations group already exists ")
		}
	}
	opGroup := ops{name: name, id: "TODO"}
	opGroup.engineers = make(map[string]engineer)
	operations[opGroup.id] = opGroup
	return opGroup, nil
}

func newEngineer(name string, email string) (engineer, error) {
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
func addEngineerTo_Op(op_id string, engineer_id string) (bool, error) {

	engineer_val, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	op_val, exists := operations[op_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	_, exists = op_val.engineers[engineer_id]
	if exists {
		return false, errors.New(" Engineer already exists inside specified Operations group ")
	}
	op_val.engineers[engineer_id] = engineer_val

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
	_, exists = dev_val.engineers[engineer_id]
	if exists {
		return false, errors.New(" Engineer already exists inside specified Operations group ")
	}
	dev_val.engineers[engineer_id] = engineer_val

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
	_, exists = devops_val.dev[dev_id]
	if exists {
		return false, errors.New(" Developer already exists inside specified Developer Operations group ")
	}
	devops_val.dev[dev_id] = dev_val

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
	_, exists = devops_val.ops[op_id]
	if exists {
		return false, errors.New(" Developer already exists inside specified Developer Operations group ")
	}
	devops_val.ops[op_id] = op_val

	return true, nil

}

// *********************************************//
// functions to delete resources from other resources//
func deleteEngineerFrom_Op(op_id string, engineer_id string) (bool, error) {

	_, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	op_val, exists := operations[op_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	_, exists = op_val.engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exists inside specified Operations group ")
	}
	delete(op_val.engineers, engineer_id)

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
	_, exists = dev_val.engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesnt exists inside specified Operations group ")
	}
	delete(dev_val.engineers, engineer_id)

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
	_, exists = devops_val.dev[dev_id]
	if !exists {
		return false, errors.New(" Developer doesn't exists inside specified Developer Operations group ")
	}
	delete(devops_val.dev, dev_id)

	return true, nil

}

func deleteOpFrom_DevOps(devops_id string, op_id string) (bool, error) {

	_, exists := developers[op_id]
	if !exists {
		return false, errors.New(" Developer group doesn't exist ")
	}
	devops_val, exists := developer_operations[devops_id]
	if !exists {
		return false, errors.New(" Developer Oerations group doesn't exist ")
	}
	_, exists = devops_val.dev[op_id]
	if !exists {
		return false, errors.New(" Developer doesn't exists inside specified Developer Operations group ")
	}
	delete(devops_val.ops, op_id)

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
	for _, value := range developer_operations {
		for key, _ := range value.dev {
			if key == dev_id {
				delete(value.dev, dev_id)
			}
		}
	}
	delete(developers, dev_id)
	return true, nil
}

func deleteOp(op_id string) (bool, error) {
	_, exists := operations[op_id]
	if !exists {
		return false, errors.New(" Operations group doesn't exist ")
	}
	for _, value := range developer_operations {
		for key, _ := range value.dev {
			if key == op_id {
				delete(value.dev, op_id)
			}
		}
	}
	delete(operations, op_id)
	return true, nil
}

func deleteEngineer(engineer_id string) (bool, error) {
	_, exists := engineers[engineer_id]
	if !exists {
		return false, errors.New(" Engineer doesn't exist ")
	}
	for _, value := range developer_operations {
		for _, value1 := range value.dev {
			for key, _ := range value1.engineers {
				if key == engineer_id {
					delete(value1.engineers, engineer_id)
					delete(developers[value1.id].engineers, engineer_id)
				}
			}
		}
	}
	for _, value := range developer_operations {
		for _, value1 := range value.ops {
			for key, _ := range value1.engineers {
				if key == engineer_id {
					delete(value1.engineers, engineer_id)
					delete(operations[value1.id].engineers, engineer_id)
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
	} else {
		return false, errors.New(" Engineer doesn't exist ")
	}
	//For updating the values for developers inside global developer_operations map
	for _, devops_val := range developer_operations {
		for _, dev_val := range devops_val.dev {
			for key, engineer_val := range dev_val.engineers {
				if key == engineer_id {
					engineer_val.Email = email
					engineer_val.Name = name
					//For updating the values inside global developer map
					dev_map_val, exists := developers[dev_val.id]
					if exists {
						dev_map_engineer_val, exists := dev_map_val.engineers[engineer_id]
						if exists {
							dev_map_engineer_val.Email = email
							dev_map_engineer_val.Name = name
						}
					}
				}
			}
		}
	}

	//For updating the values for operations inside global developer_operations map
	for _, devops_val := range developer_operations {
		for _, ops_val := range devops_val.ops {
			for key, engineer_val := range ops_val.engineers {
				if key == engineer_id {
					engineer_val.Email = email
					engineer_val.Name = name
					//For updating the values inside global operations map
					ops_map_val, exists := operations[ops_val.id]
					if exists {
						ops_map_engineer_val, exists := ops_map_val.engineers[engineer_id]
						if exists {
							ops_map_engineer_val.Email = email
							ops_map_engineer_val.Name = name
						}
					}
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
		dev_map_val.name = name
	} else {
		return false, errors.New(" Doesn't exist in the developers group")
	}
	//For updating the values for developers inside global developer_operations map
	for _, devops_val := range developer_operations {
		for key, dev_val := range devops_val.dev {
			if key == dev_id {
				dev_val.name = name
			}
		}
	}
	return true, nil
}

func updateOps(ops_id string, name string) (bool, error) {
	//For global dev map
	op_map_val, exists := operations[ops_id]
	if exists {
		op_map_val.name = name
	} else {
		return false, errors.New(" Doesn't exist in the operations group")
	}
	//For updating the values for developers inside global developer_operations map
	for _, devops_val := range developer_operations {
		for key, ops_val := range devops_val.ops {
			if key == ops_id {
				ops_val.name = name
			}
		}
	}
	return true, nil
}

func updateDevOps(devops_id string) (bool, error) {
	//TODO
	return true, nil
}

//*****************************//

func verifyEmail(email string) bool {
	result, _ := regexp.MatchString("^([a-z]|[0-9])+@[a-z]+\\.[a-z]+(\\.[a-z]+)*$", email)
	return result
}

// server GET handlers
func getEngineer(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, engineers)
}

// server POST handlers
func postEngineer(c *gin.Context) {
	var jsonData engineer    //object that gets name and email from POST request
	var curEngineer engineer //object recieved from newEngineer

	err := c.BindJSON(&jsonData)

	if err != nil {
		return
	}

	curEngineer, err = newEngineer(jsonData.Name, jsonData.Email)
	engineers[curEngineer.Id] = curEngineer
	c.IndentedJSON(http.StatusCreated, engineers)
}

func main() {
	router := gin.Default()

	//GET routes
	router.GET("/engineers", getEngineer)
	//POST routes
	router.POST("/engineers", postEngineer)

	//runs server
	router.Run("localhost:8080")
}
