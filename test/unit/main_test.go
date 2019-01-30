package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/joatmon08/hello"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnHello(t *testing.T) {
	request := httptest.NewRequest("GET", "http://myapp/hello", nil)
	writer := httptest.NewRecorder()
	hello.Hello(writer, request)
	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Response should be ok")
	assert.Equal(t, "Hello World!", string(body), "Response body did not match expected")
}

func TestShouldReturnHealth(t *testing.T) {
	request := httptest.NewRequest("GET", "http://myapp/health", nil)
	writer := httptest.NewRecorder()
	hello.Health(writer, request)
	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Response should be ok")
	assert.Equal(t, "I'm healthy!", string(body), "Response body did not match expected")
}

func TestShouldReturnErrorForPhone(t *testing.T) {
	request := httptest.NewRequest("GET", "http://myapp/phone", nil)
	writer := httptest.NewRecorder()
	hello.Phone(writer, request)
	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode, "Response should return error code")
	assert.Equal(t, "I could not connect to http://nginx!", string(body), "Response body did not match expected")
}

func TestShouldReturnSuccessForPhone(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hi!")
	}))
	defer ts.Close()

	host := strings.Replace(ts.URL, "http://", "", -1)
	expected := fmt.Sprintf("I connected to http://%s!", host)

	requestURL := fmt.Sprintf("http://myapp/phone?targetService=%s", host)

	request := httptest.NewRequest("GET", requestURL, nil)
	writer := httptest.NewRecorder()
	hello.Phone(writer, request)
	response := writer.Result()
	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode, "Response should return ok")
	assert.Equal(t, expected, string(body), "Response body did not match expected")
}

func TestShouldGenerateCPU(t *testing.T) {
	start := time.Now()
	request := httptest.NewRequest("GET", "http://myapp/cpu?testTime=5s", nil)
	writer := httptest.NewRecorder()

	hello.GenerateCPU(writer, request)
	elapsed := time.Since(start)
	elapsedSeconds := fmt.Sprintf("%.0fs", elapsed.Seconds())

	response := writer.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode, "Response should return ok")
	assert.Equal(t, elapsedSeconds, "5s", "Function didn't run with time expected")
}

func TestShouldReturnNotFound(t *testing.T) {
	request := httptest.NewRequest("GET", "http://myapp/nonexistent", nil)
	writer := httptest.NewRecorder()
	hello.NotFound(writer, request)

	response := writer.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode, "Response should return not found")
}
