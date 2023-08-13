package basicauth

import (
	"testing"
)

func TestIsAuthorizedTrue(t *testing.T) {
	input := map[string]string{"authorization": "Basic YWRtaW46cGFzc3dvcmQ="}
	expectedUsername := "admin"
	expectedPassword := "password"

	result := IsAuthorized(input, expectedUsername, expectedPassword)

	if !result {
		t.Fail()
	}
}

func TestIsAuthorizedFalse(t *testing.T) {
	input := map[string]string{"authorization": "Basic YWRtaW46cGFzc3dvcmQ="}
	expectedUsername := "admin"
	expectedPassword := "other"

	result := IsAuthorized(input, expectedUsername, expectedPassword)

	if result {
		t.Fail()
	}
}

func TestIsAuthorizedFalseInvalidHeader(t *testing.T) {
	input := map[string]string{"authorization": "Basic invalid"}
	expectedUsername := "admin"
	expectedPassword := "password"

	result := IsAuthorized(input, expectedUsername, expectedPassword)

	if result {
		t.Fail()
	}
}
