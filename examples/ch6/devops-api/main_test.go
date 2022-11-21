package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test email validation
type emailTest struct {
	description string
	email       string
	expected    bool
}

/******* structs request testing *******/
type requestEngineerTest struct {
	description  string
	testEngineer engineer
	expected     int
}

type requestDevTest struct {
	description string
	testDev     dev
	expected    int
}

type requestOpTest struct {
	description string
	testOp      ops
	expected    int
}

type requestDevOpsTest struct {
	description string
	testDev     devops
	expected    int
}

/********************************************/

var verifyEmailTests = []emailTest{
	emailTest{"simple valid email format", "bob@bob.com", true},
	emailTest{"valid email with extended domain name", "b0b123@clever.ask.who", true},
	emailTest{"special characters in username", "bob#@bob.com", false},
	emailTest{"no specificied domain name", "bob@bob", false},
	emailTest{"empty string for email", "", false},
}

/******* Slices of Test Cases for Engineer Request *******/
var verifyPostEngineer = []requestEngineerTest{
	requestEngineerTest{"contains client side Id field", engineer{Name: "Bobs Burgers", Id: getRandId(5), Email: "bob@gmail.com"}, http.StatusCreated},
	requestEngineerTest{"no client side Name field", engineer{Email: "bob@gmail.com"}, http.StatusBadRequest},
	requestEngineerTest{"no client side Email field", engineer{Name: "Bobs Burgers"}, http.StatusBadRequest},
	requestEngineerTest{"different engineer", engineer{Name: "Steven Mendez", Email: "Min3craftSt3v3@gmail.com"}, http.StatusCreated},
	requestEngineerTest{"duplicate engineer", engineer{Name: "Bobs Burgers", Email: "bob@gmail.com"}, http.StatusBadRequest},
	requestEngineerTest{"client side JSON object with empty fields", engineer{Name: "", Id: "", Email: ""}, http.StatusBadRequest},
}

var verifyPutEngineer = []requestEngineerTest{
	//Created with client side id TODO: fix where id cannot be created via client side
	requestEngineerTest{"Should update name and email of id 1 engineer", engineer{Name: "Not Bob", Id: "1", Email: "notbob@gmail.com"}, http.StatusOK},
	requestEngineerTest{"No id", engineer{Name: "Not Bob", Email: "notbob@gmail.com"}, http.StatusBadRequest},
}

var verifyDeleteEngineer = []requestEngineerTest{
	requestEngineerTest{"should delete nothing and fail", engineer{Name: "failed", Id: "40"}, http.StatusBadRequest},
	requestEngineerTest{"should delete nothing and fail since no id", engineer{Name: "NoId"}, http.StatusBadRequest},
	requestEngineerTest{"should delete test engineer and pass", engineer{Id: "5"}, http.StatusOK},
	requestEngineerTest{"duplicate this should fail", engineer{Id: "5"}, http.StatusBadRequest},
}

/********************************************/

/******* Slices of Test Cases for Developer Resource Request *******/
var verifyPostDev = []requestDevTest{
	requestDevTest{"simple developer resource creation", dev{Name: "dev_ferrets"}, http.StatusCreated},
	requestDevTest{"no client side Name field", dev{}, http.StatusBadRequest},
	requestDevTest{"different developer resource", dev{Name: "dev_bengal"}, http.StatusCreated},
	requestDevTest{"duplicate developer resource", dev{Name: "dev_ferrets"}, http.StatusBadRequest},
	requestDevTest{"client side JSON object with empty fields", dev{Name: ""}, http.StatusBadRequest},
}

var verifyPutDev = []requestDevTest{
	//Created with client side id TODO: fix where id cannot be created via client side
	requestDevTest{"should update name id 2 developer resource", dev{Name: "notferrets", Id: "2"}, http.StatusOK},
	requestDevTest{"No id", dev{Name: "dev_notferrets"}, http.StatusBadRequest},
}

var verifyDeleteDev = []requestDevTest{
	//requestDevTest{"should delete nothing and fail", dev{Name: "failed", Id: "40"}, http.StatusBadRequest},
	//requestDevTest{"should delete nothing and fail since no id", dev{Name: "NoId"}, http.StatusBadRequest},
	requestDevTest{"should delete test developer resource and pass", dev{Id: "4"}, http.StatusOK},
	requestDevTest{"duplicate this should fail", dev{Id: "4"}, http.StatusBadRequest},
}

/********************************************/

/******* Slices of Test Cases for Operations Resource Request *******/
var verifyPostOp = []requestOpTest{
	requestOpTest{"simple operation resource creation", ops{Name: "op_ferrets"}, http.StatusCreated},
	requestOpTest{"no client side Name field", ops{}, http.StatusBadRequest},
	requestOpTest{"different operation resource", ops{Name: "op_bengal"}, http.StatusCreated},
	requestOpTest{"duplicate operation resource", ops{Name: "op_ferrets"}, http.StatusBadRequest},
	requestOpTest{"client side JSON object with empty fields", ops{Name: ""}, http.StatusBadRequest},
}

var verifyPutOp = []requestOpTest{
	//Created with client side id TODO: fix where id cannot be created via client side
	requestOpTest{"should update name id 2 operation resource", ops{Name: "op_notferrets", Id: "2"}, http.StatusOK},
	requestOpTest{"No id", ops{Name: "op_notferrets"}, http.StatusBadRequest},
}

var verifyDeleteOp = []requestOpTest{
	requestOpTest{"should delete nothing and fail", ops{Name: "failed", Id: "40"}, http.StatusBadRequest},
	requestOpTest{"should delete nothing and fail since no id", ops{Name: "NoId"}, http.StatusBadRequest},
	requestOpTest{"should delete test operation resource and pass", ops{Id: "2"}, http.StatusOK},
	requestOpTest{"duplicate this should fail", ops{Id: "2"}, http.StatusBadRequest},
}

/********************************************/

func TestVerifyEmail(t *testing.T) {
	for _, test := range verifyEmailTests {
		result := verifyEmail(test.email)
		if result != test.expected {
			t.Errorf("\nTest: %s\nEmail: %s\nExpected: %t, Received: %t", test.description, test.email, test.expected, result)
		}
	}
}

/******* Engineer Mock Requests *******/
func mockJsonPostEngineer(c *gin.Context, content engineer) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func mockJsonPutEngineer(c *gin.Context, content engineer) {
	c.Request.Method = "PUT"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{gin.Param{Key: "id", Value: content.Id}}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func mockJsonDeleteEngineer(c *gin.Context, content engineer) {
	c.Request.Method = "DELETE"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{gin.Param{Key: "id", Value: content.Id}}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

/********************************************/

/******* Dev Mock Requests *******/
func mockJsonPostDev(c *gin.Context, content dev) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func mockJsonPutDev(c *gin.Context, content dev) {
	c.Request.Method = "PUT"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{gin.Param{Key: "id", Value: content.Id}}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func mockJsonDeleteDev(c *gin.Context, content dev) {
	c.Request.Method = "DELETE"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{gin.Param{Key: "id", Value: content.Id}}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

/********************************************/

/******* Op Mock Requests *******/
func mockJsonPostOp(c *gin.Context, content ops) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func mockJsonPutOp(c *gin.Context, content ops) {
	c.Request.Method = "PUT"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{gin.Param{Key: "id", Value: content.Id}}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func mockJsonDeleteOp(c *gin.Context, content ops) {
	c.Request.Method = "DELETE"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{gin.Param{Key: "id", Value: content.Id}}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

/********************************************/

/******* Test Runs for Engineer Requests *******/
func TestPostEngineer(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context

	for _, test := range verifyPostEngineer {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonPostEngineer(c, test.testEngineer)
		postEngineer(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

func TestPutEngineer(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	engineers = append(engineers, &engineer{Name: "Bob", Id: "1", Email: "bob@gmail.com"})

	for _, test := range verifyPutEngineer {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonPutEngineer(c, test.testEngineer)
		putEngineer(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

func TestDeleteRequestEngineer(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	engineers = append(engineers, &engineer{Name: "Bob", Id: "5", Email: "bob@gmail.com"})

	for _, test := range verifyDeleteEngineer {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonDeleteEngineer(c, test.testEngineer)
		deleteRequestEngineer(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

/********************************************/

/******* Test Runs for Developer Resource Requests *******/
func TestPostDev(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context

	for _, test := range verifyPostDev {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonPostDev(c, test.testDev)
		postDev(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

func TestPutDev(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	developers = append(developers, &dev{Name: "ferrets", Id: "2"})

	for _, test := range verifyPutDev {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonPutDev(c, test.testDev)
		putDev(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

func TestDeleteRequestDev(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	developers = append(developers, &dev{Name: "ferrets", Id: "4"})

	for _, test := range verifyDeleteDev {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonDeleteDev(c, test.testDev)
		deleteRequestDev(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
		fmt.Println(len(developers))
		fmt.Println(developers[3])
	}
}

/********************************************/

/******* Test Runs for Operation Resource Requests *******/
func TestPostOp(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context

	for _, test := range verifyPostOp {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonPostOp(c, test.testOp)
		postOp(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

func TestPutOp(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	operations = append(operations, &ops{Name: "ferrets", Id: "2"})

	for _, test := range verifyPutOp {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonPutOp(c, test.testOp)
		putOp(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

func TestDeleteRequestOp(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context
	operations = append(operations, &ops{Name: "ferrets", Id: "2"})

	for _, test := range verifyDeleteOp {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonDeleteOp(c, test.testOp)
		deleteRequestOp(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

/********************************************/

func TestNewEngineer(t *testing.T) {
	result, err := newEngineer("test_engineer", "test@gmail.com")
	if result.Email != "test@gmail.com" || result.Name != "test_engineer" {
		t.Errorf("Expected name %s and email %s were not returned", result.Name, result.Email)
	}
	if err != nil {
		t.Errorf("Expected no errors but one was returned")
	}
}

func TestNewEngineerBadName(t *testing.T) {
	result, err := newEngineer("", "test@gmail.com")
	if err == nil {
		t.Errorf("Expected name %s was not returned", result.Name)
	}
	if result != (&engineer{}) {
		t.Errorf("Expected empty struct but one was not returned")
	}

}

func TestNewDev(t *testing.T) {

	result, err := newDev(dev{Name: "test_devs"})
	if result.Name != "test_devs" {
		t.Errorf("Expected name %s was not returned", result.Name)
	}
	if err != nil {
		t.Errorf("Expected no errors but one was returned")
	}
}
func TestNewDevBadName(t *testing.T) {
	_, err := newDev(dev{Name: ""})
	if err == nil {
		t.Errorf("Expected name error but one was not returned")
	}
}

func TestNewOp(t *testing.T) {
	result, err := newOp(ops{Name: "test_ops"})
	if result.Name != "test_ops" {
		t.Errorf("Expected name %s was not returned", result.Name)
	}
	if err != nil {
		t.Errorf("Expected no errors but one was returned")
	}
}

func TestNewDevOps(t *testing.T) {
	_, err := newDevOps(devops{})
	if err != nil {
		t.Errorf("Expected no errors but one was returned")
	}
	//possibly test if maps are empty/initialized
}

func TestDeleteEngineer(t *testing.T) {
	result, _ := newEngineer("test_engineer2", "test@gmail.com")
	delete_result, _ := deleteEngineer(result.Id)
	if delete_result == false {
		t.Errorf("Expected true to be returned but %t was instead'", delete_result)
	}
	_, delete2_error := deleteEngineer(result.Id)
	if delete2_error == nil {
		t.Errorf("Expected error to occur on second delete of same object but nil was returned")
	}
}

/*
func TestDeleteDev(t *testing.T) {

}

func TestDeleteOps(t *testing.T) {

}

func TestDeleteDevOps(t *testing.T) {

}


func TestUpdateEngineer(t *testing.T) {

}

func TestUpdateDev(t *testing.T) {

}
func TestUpdateOps(t *testing.T) {

}

func TestAddOpsTo_DevOps(t *testing.T) {

}
func TestAddDevTo_DevOps(t *testing.T) {

}
func TestAddEngineerTo_Dev(t *testing.T) {

}
func TestAddEngineerTo_Ops(t *testing.T) {

}

func TestDeleteOpsFrom_DevOps(t *testing.T) {

}
func TestDeleteDevFrom_DevOps(t *testing.T) {

}
func TestDeleteEngineerFrom_Dev(t *testing.T) {

}
func TestDeleteEngineerFrom_Ops(t *testing.T) {

}
*/
