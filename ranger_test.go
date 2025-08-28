package ranger

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetFakeRangerServer() *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		policy := `[
			{
				"id": 1,
				"name": "test policy",
				"service": "kafka"
			},
			{
				"id": 2,
				"name": "another policy",
				"service": "hive"
			}
		]`

		kafkaPolicy := `[
			{
				"id": 1,
				"name": "test policy",
				"service": "kafka"
			}
		]`

		hivePolicy := `[
			{
				"id": 2,
				"name": "another policy",
				"service": "hive"
			}
		]`

		switch r.URL.Path {
		case "/":
			w.WriteHeader(http.StatusOK)
		case "/service/public/v2/api/policy":
			switch r.Method {
			case http.MethodGet:
				w.WriteHeader(http.StatusOK)
				if r.URL.Query().Get("serviceName") == "kafka" {
					if _, err := w.Write([]byte(kafkaPolicy)); err != nil {
						fmt.Printf("Error writing response: %v", err)
					}
				} else if r.URL.Query().Get("serviceName") == "hive" {
					if _, err := w.Write([]byte(hivePolicy)); err != nil {
						fmt.Printf("Error writing response: %v", err)
					}
				} else {
					if _, err := w.Write([]byte(policy)); err != nil {
						fmt.Printf("Error writing response: %v", err)
					}
				}
			case http.MethodPost:
				// The request contains a policy in the body
				// Read the body and return a created policy
				var newPolicy Policy
				err := json.NewDecoder(r.Body).Decode(&newPolicy)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					if _, writeErr := fmt.Fprintf(w, "Bad Request: %v", err); writeErr != nil {
						fmt.Printf("Error writing response: %v", writeErr)
					}
					return
				}

				createdPolicy := Policy{
					ID:      1,
					Name:    newPolicy.Name,
					Service: newPolicy.Service,
				}

				jsonBytes, err := json.Marshal(createdPolicy)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					if _, err := w.Write([]byte("Internal Server Error")); err != nil {
						fmt.Printf("Error writing response: %v", err)
					}
					return
				}
				if _, err := w.Write(jsonBytes); err != nil {
					fmt.Printf("Error writing response: %v", err)
				}
			}
		// case with a specific policy number
		case "/service/public/v2/api/policy/1":
			switch r.Method {
			case http.MethodPut:
				// The request contains a policy in the body to update
				var policyToUpdate Policy
				err := json.NewDecoder(r.Body).Decode(&policyToUpdate)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					if _, writeErr := fmt.Fprintf(w, "Bad Request: %v", err); writeErr != nil {
						fmt.Printf("Error writing response: %v", writeErr)
					}
					return
				}

				updatedPolicy := Policy{
					ID:      policyToUpdate.ID,
					Name:    policyToUpdate.Name,
					Service: policyToUpdate.Service,
				}

				jsonBytes, err := json.Marshal(updatedPolicy)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					if _, err := w.Write([]byte("Internal Server Error")); err != nil {
						fmt.Printf("Error writing response: %v", err)
					}
					return
				}
				w.WriteHeader(http.StatusOK)
				if _, err := w.Write(jsonBytes); err != nil {
					fmt.Printf("Error writing response: %v", err)
				}
			case http.MethodDelete:
				w.WriteHeader(http.StatusNoContent)
			}
		default:
			w.WriteHeader(http.StatusNotFound)
			if _, err := w.Write([]byte("Not Found")); err != nil {
				fmt.Printf("Error writing response: %v", err)
			}
		}
	}))

	return server
}

func TestNewClient(t *testing.T) {
	uri := "https://example.ranger.com"
	username := "testuser"
	password := "testpassword"

	c := NewClient(uri, username, password)

	if c.BaseURL != uri {
		t.Errorf("expected BaseURL %s, got %s", uri, c.BaseURL)
	}

	if c.Username != username {
		t.Errorf("expected Username %s, got %s", username, c.Username)
	}

	if c.Password != password {
		t.Errorf("expected Password %s, got %s", password, c.Password)
	}

	if c == nil {
		t.Error("expected client to be created, got nil")
	}
}

func TestDoRequest(t *testing.T) {
	testServer := GetFakeRangerServer()

	c := NewClient(testServer.URL, "testuser", "testpassword")

	req, err := http.NewRequest("GET", testServer.URL, nil)

	if err != nil {
		t.Errorf("expected no error creating request, got %v", err)
	}

	resp, err := c.doRequest(req)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if resp == nil {
		t.Error("expected response, got nil")
		return
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestGetPolicies(t *testing.T) {
	testServer := GetFakeRangerServer()

	defer testServer.Close()

	c := NewClient(testServer.URL, "testuser", "testpassword")

	policies, err := c.GetPolicies()

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(policies) != 2 {
		t.Errorf("expected 2 policies, got %d", len(policies))
	}

	if policies[0].ID != 1 {
		t.Errorf("expected policy ID 1, got %d", policies[0].ID)
	}

	if policies[0].Name != "test policy" {
		t.Errorf("expected policy name 'test policy', got '%s'", policies[0].Name)
	}

	if policies[0].Service != "kafka" {
		t.Errorf("expected policy service 'kafka', got '%s'", policies[0].Service)
	}

	if policies[1].ID != 2 {
		t.Errorf("expected policy ID 2, got %d", policies[1].ID)
	}

	if policies[1].Name != "another policy" {
		t.Errorf("expected policy name 'another policy', got '%s'", policies[1].Name)
	}

	if policies[1].Service != "hive" {
		t.Errorf("expected policy service 'hive', got '%s'", policies[1].Service)
	}
}

func TestGetPoliciesWithService(t *testing.T) {
	testServer := GetFakeRangerServer()

	defer testServer.Close()

	c := NewClient(testServer.URL, "testuser", "testpassword")

	policies, err := c.GetPolicies("hive")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(policies) != 1 {
		t.Errorf("expected 1 policy for service 'hive', got %d", len(policies))
	}

	if policies[0].ID != 2 {
		t.Errorf("expected policy ID 2, got %d", policies[0].ID)
	}

	if policies[0].Name != "another policy" {
		t.Errorf("expected policy name 'another policy', got '%s'", policies[0].Name)
	}

	if policies[0].Service != "hive" {
		t.Errorf("expected policy service 'hive', got '%s'", policies[0].Service)
	}
}

func TestCreatePolicy(t *testing.T) {
	testServer := GetFakeRangerServer()

	defer testServer.Close()

	c := NewClient(testServer.URL, "testuser", "testpassword")

	newPolicy := Policy{
		Name:    "New policy",
		Service: "kafka",
	}

	createdPolicy, err := c.CreatePolicy(&newPolicy)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if createdPolicy == nil {
		t.Error("expected created policy, got nil")
		return
	}

	if createdPolicy.ID == 0 {
		t.Error("expected created policy to have a non-zero ID")
	}

	if createdPolicy.Name != newPolicy.Name {
		t.Errorf("expected policy name '%s', got '%s'", newPolicy.Name, createdPolicy.Name)
	}

	if createdPolicy.Service != newPolicy.Service {
		t.Errorf("expected policy service '%s', got '%s'", newPolicy.Service, createdPolicy.Service)
	}
}

func TestUpdatePolicy(t *testing.T) {
	testServer := GetFakeRangerServer()

	defer testServer.Close()

	c := NewClient(testServer.URL, "testuser", "testpassword")

	policyToUpdate := Policy{
		ID:      1,
		Name:    "New policy",
		Service: "kafka",
	}

	updatedPolicy, err := c.UpdatePolicy(&policyToUpdate)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if updatedPolicy == nil {
		t.Error("expected created policy, got nil")
		return
	}

	if updatedPolicy.ID != policyToUpdate.ID {
		t.Errorf("expected created policy to have ID %d, got %d", policyToUpdate.ID, updatedPolicy.ID)
	}

	if updatedPolicy.Name != policyToUpdate.Name {
		t.Errorf("expected policy name '%s', got '%s'", policyToUpdate.Name, updatedPolicy.Name)
	}

	if updatedPolicy.Service != policyToUpdate.Service {
		t.Errorf("expected policy service '%s', got '%s'", policyToUpdate.Service, updatedPolicy.Service)
	}
}

func TestDeletePolicy(t *testing.T) {
	testServer := GetFakeRangerServer()

	defer testServer.Close()

	c := NewClient(testServer.URL, "testuser", "testpassword")

	policyToDelete := Policy{
		ID: 1,
	}

	err := c.DeletePolicy(policyToDelete.ID)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
