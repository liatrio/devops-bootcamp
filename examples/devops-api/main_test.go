package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// test email validation
type emailTest struct {
	email    string
	expected bool
}

type postEngineerTest struct {
	testEngineer engineer
	expected     int
}

var verifyEmailTests = []emailTest{
	emailTest{"bob@bob.com", true},
	emailTest{"b0b123@clever.ask.who", true},
	emailTest{"bob#@bob.com", false},
	emailTest{"bob@bob", false},
}

var verifyPostEngineer = []postEngineerTest{
	//regular creation
	postEngineerTest{engineer{Name: "Bobs Burgers", Id: getRandId(5), Email: "bob@gmail.com"}, http.StatusCreated},
	//no name
	postEngineerTest{engineer{Id: getRandId(5), Email: "bob@gmail.com"}, http.StatusBadRequest},
	//no email
	postEngineerTest{engineer{Name: "Bobs Burgers", Id: getRandId(5)}, http.StatusBadRequest},
	//no id
	postEngineerTest{engineer{Name: "Bobs Burgers", Email: "bob@gmail.com"}, http.StatusCreated},
}

func TestVerifyEmail(t *testing.T) {
	for _, test := range verifyEmailTests {
		result := verifyEmail(test.email)
		if result != test.expected {
			t.Errorf("Email: %s\nExpected: %t, Received: %t", test.email, test.expected, result)
		}
	}
}

func MockJsonPost(c *gin.Context, content engineer) {
	c.Request.Method = "POST"
	c.Request.Header.Set("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(content)
	if err != nil {
		panic(err)
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBytes))
}

func TestPostEngineer(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	var w *httptest.ResponseRecorder
	var c *gin.Context

	for _, test := range verifyPostEngineer {
		w = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(w)
		c.Request = &http.Request{
			Header: make(http.Header),
		}
		MockJsonPost(c, test.testEngineer)
		postEngineer(c)
		if test.expected != w.Code {
			t.Errorf("Expected: Status Code %d, Received: Status Code %d", test.expected, w.Code)
		}
	}
}
