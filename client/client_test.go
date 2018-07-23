package client

import (
	"errors"
	"net/http"
	"testing"

	"github.com/ashwanthkumar/marathonctl/util"
	"github.com/parnurzeal/gorequest"
	"github.com/stretchr/testify/assert"
)

type Error struct {
	Message string `json:"message"`
}

func TestHandleOn200Success(t *testing.T) {
	body, err := handle(createMockResponse(200), "body", []error{})
	assert.NoError(t, err)
	assert.Equal(t, body, "body")
}

func TestHandleOn201_WhenNewAppIsCreated(t *testing.T) {
	body, err := handle(createMockResponse(201), "body", []error{})
	assert.NoError(t, err)
	assert.Equal(t, body, "body")
}

func TestHandleOnNon200Or201(t *testing.T) {
	body, err := handle(createMockResponse(404), createErrorResponse("App '/foo' not found"), []error{errors.New("NotFound")})

	expectedBody := "{\n  \"message\": \"App '/foo' not found\"\n}"
	assert.Equal(t, body, expectedBody)

	expectedError := errors.New("Error(s): NotFound App '/foo' not found")
	assert.Equal(t, err, expectedError)
}

func TestHandleOnNon200Or201WhenBodyIsEmpty(t *testing.T) {
	mockResponse := createMockResponse(404)
	mockResponse.Status = "(404) Resource Not found"
	body, err := handle(mockResponse, "", []error{})

	// empty in empty out
	assert.Equal(t, body, "")

	expectedError := errors.New("(404) Resource Not found")
	assert.Equal(t, err, expectedError)
}

func TestHandleOn422UnprocessableEntity(t *testing.T) {
	body, err := handle(createMockResponse(422), createErrorResponseNew("Object is not valid"), []error{errors.New("UnprocessableEntity")})

	expectedBody := "{\"message\": \"Object is not valid\", \"details\": [{\"path\": \"/healthChecks(0)\",\"errors\": [\"Health check port indices must address an element of the ports array or container port mappings.\"]}]}"
	assert.Equal(t, body, expectedBody)

	expectedError := errors.New("Error(s): UnprocessableEntity Health check port indices must address an element of the ports array or container port mappings.")
	assert.Equal(t, err, expectedError)
}

func TestHandleNilResponse(t *testing.T) {
	_, err := handle(nil, createErrorResponse("Invalid Request"), []error{errors.New("InvalidRequest")})
	expectedError := errors.New("InvalidRequest")
	assert.Equal(t, err, expectedError)
}

func TestCombineErrors(t *testing.T) {
	input := []error{
		errors.New("error1"),
		errors.New("error2"),
	}
	err := combineErrors(input)
	expectedError := errors.New("Error(s): error1 error2")
	assert.Equal(t, err, expectedError)
}

func createMockResponse(statusCode int) gorequest.Response {
	httpResponse := &http.Response{
		StatusCode: statusCode,
	}

	var response gorequest.Response
	response = httpResponse
	return response
}

func createErrorResponse(message string) string {
	err := Error{
		Message: message,
	}
	return util.ToJson(err)
}

func createErrorResponseNew(message string) string {
	_ = message
	return "{\"message\": \"Object is not valid\", \"details\": [{\"path\": \"/healthChecks(0)\",\"errors\": [\"Health check port indices must address an element of the ports array or container port mappings.\"]}]}"
}
