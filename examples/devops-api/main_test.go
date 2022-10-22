package main

import "testing"

type emailTest struct {
	email    string
	expected bool
}

var verifyEmailTests = []emailTest{
	emailTest{"bob@bob.com", true},
}

func testVerifyEmail(t *testing.T) {
	for _, test := range verifyEmailTests {
		result := verifyEmail(test.email)
		if result != test.expected {
			t.Errorf("Expected: %t, Received: %t", test.expected, result)
		}
	}
}
