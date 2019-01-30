package test

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	url string
)

func setup() error {
	var isPresent bool
	url, isPresent = os.LookupEnv("HELLO_ENDPOINT")
	if !isPresent {
		return errors.New("Please define HELLO_ENDPOINT")
	}
	return nil
}

func TestMain(m *testing.M) {
	if err := setup(); err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestShouldReturnHello(t *testing.T) {
	endpoint := fmt.Sprintf("%s/hello", url)
	request, _ := http.NewRequest("GET", endpoint, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Could not execute request. %s", err)
	}
	body, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Response should be ok")
	assert.Equal(t, "Hello World!", string(body), "Response body did not match expected")
}

func TestShouldReturnErrorForPhone(t *testing.T) {
	endpoint := fmt.Sprintf("%s/phone", url)
	request, _ := http.NewRequest("GET", endpoint, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Could not execute request. %s", err)
	}
	body, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusInternalServerError, response.StatusCode, "Response should return error code")
	assert.Equal(t, "I could not connect to http://nginx!", string(body), "Response body did not match expected")
}

func TestShouldReturnSuccessForPhone(t *testing.T) {
	host := "google.com"
	expected := fmt.Sprintf("I connected to http://%s!", host)
	endpoint := fmt.Sprintf("%s/phone?targetService=%s", url, host)
	request, _ := http.NewRequest("GET", endpoint, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Could not execute request. %s", err)
	}
	body, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusOK, response.StatusCode, "Response should return ok")
	assert.Equal(t, expected, string(body), "Response body did not match expected")
}
func TestShouldReturnNotFound(t *testing.T) {
	endpoint := fmt.Sprintf("%s/notanendpoint", url)
	request, _ := http.NewRequest("GET", endpoint, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		t.Errorf("Could not execute request. %s", err)
	}

	assert.Equal(t, http.StatusNotFound, response.StatusCode, "Response should return not found")
}
