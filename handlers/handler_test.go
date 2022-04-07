package handlers

import (
	constants "assignment_2"
	"assignment_2/structs"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const URL = "http://localhost:8080"

func TestCasesHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(CasesHandler))

	defer server.Close()

	client := http.Client{}

	res, err := client.Get(server.URL + constants.CASES_PATH)
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)

	check := "test"
	jsonString, err := json.Marshal(check)
	if err != nil {
		t.Fatal("Error when marshaling json: ", err.Error())
	}

	req, err := http.NewRequest("POST", (server.URL + constants.POLICY_PATH), bytes.NewReader(jsonString))
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	res1, err := client.Do(req)
	if err != nil {
		t.Fatal("Error when sending POST request: ", err.Error())
	}

	assert.Equal(t, res1.StatusCode, http.StatusOK)
}

func TestCasesGetRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(CasesHandler))

	defer server.Close()

	client := http.Client{}

	res, err := client.Get(server.URL + constants.CASES_PATH + "Norway")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)

	res1, err := client.Get(server.URL + constants.CASES_PATH + "Norw")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	assert.Equal(t, res1.StatusCode, http.StatusOK)
}

func TestDefaultHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(DefaultHandler))

	defer server.Close()

	client := http.Client{}

	res, err := client.Get(server.URL + constants.DEFAULT_PATH)
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	assert.Equal(t, res.StatusCode, http.StatusAccepted)
}

func TestPolicyHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(PolicyHandler))

	defer server.Close()

	check := "test"
	jsonString, err := json.Marshal(check)
	if err != nil {
		t.Fatal("Error when marshaling json: ", err.Error())
	}

	client := http.Client{}

	req, err := http.NewRequest("POST", (server.URL + constants.POLICY_PATH), bytes.NewReader(jsonString))
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal("Error when sending POST request: ", err.Error())
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestGetPolicyRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(PolicyHandler))

	defer server.Close()

	client := http.Client{}

	res, err := client.Get(server.URL + constants.POLICY_PATH + "NOR/2021-01-01")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	policies := []structs.Policy{}
	err2 := json.NewDecoder(res.Body).Decode(&policies)
	if err2 != nil {
		t.Fatal("Error during decoding:", err2.Error())
	}

	assert.Equal(t, len(policies), 1)

	res1, err := client.Get(server.URL + constants.POLICY_PATH + "NOR")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}
	assert.Equal(t, res1.StatusCode, http.StatusNonAuthoritativeInfo)
}

func TestStatusHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(StatusHandler))

	defer server.Close()

	check := "test"
	jsonString, err := json.Marshal(check)
	if err != nil {
		t.Fatal("Error when marshaling json: ", err.Error())
	}

	client := http.Client{}

	req, err := http.NewRequest("POST", (server.URL + constants.STATUS_PATH), bytes.NewReader(jsonString))
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	res, err := client.Do(req)
	if err != nil {
		t.Fatal("Error when sending POST request: ", err.Error())
	}

	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func TestStatusHandlerGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(StatusHandler))

	defer server.Close()

	client := http.Client{}

	res, err := client.Get(server.URL + constants.STATUS_PATH)
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}

	statuses := []structs.Status{}
	status := structs.Status{}
	err2 := json.NewDecoder(res.Body).Decode(&status)
	if err2 != nil {
		t.Fatal("Error during decoding:", err2.Error())
	}
	statuses = append(statuses, status)

	assert.Equal(t, len(statuses), 1)
}

func aTestNotificationHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(NotificationHandler))

	defer server.Close()

	client := http.Client{}

	res, err := client.Get(server.URL + constants.NOTIFICATION_PATH + "QnzWFw6MoSbejTI8R4GO")
	if err != nil {
		t.Fatal("Get request to URL failed:", err.Error())
	}
	assert.Equal(t, res.StatusCode, http.StatusOK)
}

func aTestFirebaseMock(t *testing.T) {
	webhooks := FirebaseMock()

	assert.Equal(t, len(webhooks), 1)
}
