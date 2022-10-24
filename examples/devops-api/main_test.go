package main

import "testing"

type emailTest struct {
	email    string
	expected bool
}

var verifyEmailTests = []emailTest{
	emailTest{"bob@bob.com", true},
	emailTest{"b0b123@clever.ask.who", true},
	emailTest{"bob#@bob.com", false},
	emailTest{"bob@bob", false},
}

func TestVerifyEmail(t *testing.T) {
	for _, test := range verifyEmailTests {
		result := verifyEmail(test.email)
		if result != test.expected {
			t.Errorf("Email: %s\nExpected: %t, Received: %t", test.email, test.expected, result)
		}
	}
}
