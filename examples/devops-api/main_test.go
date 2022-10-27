package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// test email validation
type emailTest struct {
	description string
	email       string
	expected    bool
}

// test POST request for new engineer
type postEngineerTest struct {
	description  string
	testEngineer engineer
	expected     int
}

type putEngineerTest struct {
	description  string
	testEngineer engineer
	expected     int
}

var verifyEmailTests = []emailTest{
	emailTest{"simple valid email format", "bob@bob.com", true},
	emailTest{"valid email with extended domain name", "b0b123@clever.ask.who", true},
	emailTest{"special characters in username", "bob#@bob.com", false},
	emailTest{"no specificied domain name", "bob@bob", false},
	emailTest{"empty string for email", "", false},
}

var verifyPostEngineer = []postEngineerTest{
	postEngineerTest{"contains client side Id field", engineer{Name: "Bobs Burgers", Id: getRandId(5), Email: "bob@gmail.com"}, http.StatusCreated},
	postEngineerTest{"no client side Name field", engineer{Email: "bob@gmail.com"}, http.StatusBadRequest},
	postEngineerTest{"no client side Email field", engineer{Name: "Bobs Burgers"}, http.StatusBadRequest},
	postEngineerTest{"different engineer", engineer{Name: "Steven Mendez", Email: "Min3craftSt3v3@gmail.com"}, http.StatusCreated},
	postEngineerTest{"duplicate engineer", engineer{Name: "Bobs Burgers", Email: "bob@gmail.com"}, http.StatusBadRequest},
	postEngineerTest{"client side JSON object with empty fields", engineer{Name: "", Id: "", Email: ""}, http.StatusBadRequest},
}

var verifyPutEngineer = []putEngineerTest{
	//Created with client side id TODO: fix where id cannot be created via client side
	putEngineerTest{"Should update name and email of id 5 engineer", engineer{Name: "Not Bob", Id: "5", Email: "notbob@gmail.com"}, http.StatusOK},
	putEngineerTest{"No id", engineer{Name: "Not Bob", Email: "notbob@gmail.com"}, http.StatusBadRequest},
}

func TestVerifyEmail(t *testing.T) {
	for _, test := range verifyEmailTests {
		result := verifyEmail(test.email)
		if result != test.expected {
			t.Errorf("\nTest: %s\nEmail: %s\nExpected: %t, Received: %t", test.description, test.email, test.expected, result)
		}
	}
}

func mockJsonPost(c *gin.Context, content engineer) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")
	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func mockJsonPut(c *gin.Context, content engineer) {
	c.Request.Method = "PUT"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = []gin.Param{gin.Param{Key: "id", Value: content.Id}}

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func TestPostEngineer(t *testing.T) {
	var w *httptest.ResponseRecorder
	var c *gin.Context

	for _, test := range verifyPostEngineer {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonPost(c, test.testEngineer)
		postEngineer(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

func TestPutEngineer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	var w *httptest.ResponseRecorder
	var c *gin.Context
	engineers["5"] = engineer{Name: "Bob", Id: "5", Email: "bob@gmail.com"}

	for _, test := range verifyPutEngineer {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		mockJsonPut(c, test.testEngineer)
		putEngineer(c)
		if test.expected != w.Code {
			t.Errorf("\nTest: %s\nExpected: Status Code %d, Received: Status Code %d", test.description, test.expected, w.Code)
		}
	}
}

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
	if result != (engineer{}) {
		t.Errorf("Expected empty struct but one was not returned")
	}

}

func TestNewDev(t *testing.T) {
	result, err := newDev("test_devs")
	if result.Name != "test" {
		t.Errorf("Expected name %s was not returned", result.Name)
	}
	if err != nil {
		t.Errorf("Expected no errors but one was returned")
	}
}
func TestNewDevBadName(t *testing.T) {
	_, err := newDev("")
	if err == nil {
		t.Errorf("Expected name error but one was not returned")
	}
}
func TestNewOp(t *testing.T) {
	result, err := newOp("test_ops")
	if result.Name != "test_ops" {
		t.Errorf("Expected name %s was not returned", result.Name)
	}
	if err != nil {
		t.Errorf("Expected no errors but one was returned")
	}
}
func TestNewDevOps(t *testing.T) {
	_, err := newDevOps()
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
